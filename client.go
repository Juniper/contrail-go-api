//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"strconv"
	"unicode"
)

// TypeMap is used to inject the auto-generated types library.
//
// Types are generated from the OpenContrail schema and allow the library
// to operate in terms of go structs that contain fields that represent
// IF-MAP properties (metadata associated with a single Identifier) and
// arrays of references to other Identifiers (with optional metadata).
// Each auto-generated type implements the IObject interface.
type TypeMap map[string]reflect.Type

// ClientInterface defines the interface used by the contrail API client.
//
// The ObjectBase type includes a reference to this interface.
type ClientInterface interface {
	GetHttpClient() *http.Client
	GetField(IObject, string) error
	UpdateReference(*ReferenceUpdateMsg) error
}

// A client of the OpenContrail API server.
type Client struct {
	server string
	port   int
	httpClient *http.Client
}

// The Client List API returns an array of ListResult entries.
type ListResult struct {
	Fq_name []string
	Href string
	Uuid string
}

var (
	typeMap TypeMap
)

// Allocates and initialized a client.
//
// The typeMap parameter specifies a map of name, reflection Type values
// use to deserialize the data received from the server.
func NewClient(server string, port int) *Client {
	client := new(Client)
	client.server = server
	client.port = port
	client.httpClient = &http.Client{}
	return client
}

func typename(ptr IObject) string {
	name := reflect.TypeOf(ptr).Elem().Name()
	var buf []rune
	for i, c := range name {
		if unicode.IsUpper(c) {
			if i > 0 {
				buf = append(buf, '-')
			}
			buf = append(buf, unicode.ToLower(c))
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}

func (c *Client) GetHttpClient() *http.Client {
	return c.httpClient
}

// Create an object in the OpenContrail API server.
//
// The object must have been initialized with a name.
func (c *Client) Create(ptr IObject) error {
	xtype := typename(ptr)
	url := fmt.Sprintf("http://%s:%d/%ss", c.server, c.port, xtype)

	objJson, err := json.Marshal(ptr)
	if err != nil {
		return err
	}

	var rawJson json.RawMessage = objJson
	msg := map[string]*json.RawMessage {
		xtype: &rawJson,
	}
	data, err := json.Marshal(msg)

	resp, err := c.httpClient.Post(url, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	ptr.SetClient(c)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var m map[string]json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}

	return json.Unmarshal(m[xtype], ptr)
}

// Read an object from the API server.
//
// This method retrieves the object properties but not its references to
// other objects.
func (c *Client) readObject(typename string, href string) (IObject, error) {
	url := fmt.Sprintf("%s?exclude_back_refs=true&exclude_children=true",
		href)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	content, ok := m[typename]
	if !ok {
		return nil, fmt.Errorf("No %s in Response", typename)
	}

	var xtype reflect.Type = typeMap[typename]
	valueT := reflect.New(xtype)
	obj := valueT.Interface().(IObject)
	err = json.Unmarshal(*content, obj)
	if err != nil {
		return nil, err
	}
	obj.SetClient(c)
	return obj, err
}

// Given a ListResult, retrieve an object from the API server.
func (c *Client) ReadListResult(
	typename string, result *ListResult) (IObject, error) {
	return c.readObject(typename, result.Href)
}

// Given a link reference, retrieve an object from the API server.
func (c *Client) ReadReference(
	typename string, ref *Reference) (IObject, error) {
	return c.readObject(typename, ref.Href)
}

// Update the API server with the changes made in the local representation
// of the object.
//
// There is currently no mechanism to guarantee that the object as not
// been concurrently modified in the API server.
// Updates modify properties that have been marked as modified in the local
// representation.
func (c *Client) Update(ptr IObject) error {
	objJson, err := ptr.UpdateObject()
	if err != nil {
		return err
	}
	var rawJson json.RawMessage = objJson
	msg := map[string]*json.RawMessage {
		ptr.GetType(): &rawJson,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", ptr.GetHref(), bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = ptr.UpdateReferences()
	if err != nil {
		return err
	}
	ptr.UpdateDone()

	return nil
}

// Delete an object from the API server.
func (c *Client) Delete(ptr IObject) error {
	req, err := http.NewRequest("DELETE", ptr.GetHref(), nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}

// Read an object identified by UUID.
func (c *Client) FindByUuid(typename string, uuid string) (IObject, error) {
	url := fmt.Sprintf("http://%s:%d/%s/%s", c.server, c.port,
		typename, uuid)
	return c.readObject(typename, url)
}

func (c *Client) UuidByName(typename string, fqn string) (string, error) {
	url := fmt.Sprintf("http://%s:%d/fqname-to-id", c.server, c.port)
	request := struct {
		Typename string `json:"type"`
		Fq_name []string `json:"fq_name"`
	}{
		typename,
		strings.Split(fqn, ":"),
	}
	data, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
 	resp, err := c.httpClient.Post(url, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	m := struct {
		Uuid string
	}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return "", err
	}

	return m.Uuid, nil
}

// Read an object identified by fully-qualified name represented as a
// string.
func (c *Client) FindByName(typename string, fqn string) (IObject, error) {
	uuid, err := c.UuidByName(typename, fqn)
	if err != nil {
		return nil, err
	}
	href := fmt.Sprintf(
		"http://%s:%d/%s/%s", c.server, c.port, typename, uuid)
	return c.readObject(typename, href)
}

// Retrieve the list of all elements of a specific type.
func (c *Client) ListByParent(
	typename string, parent_id string, count int) ([]ListResult, error) {
	var values url.Values
	values = make(url.Values, 0)
	if len(parent_id) > 0 {
		values.Add("parent_id", parent_id)
	}
	if count > 0 {
		values.Add("count", strconv.Itoa(count))
	}

	url := fmt.Sprintf("http://%s:%d/%ss", c.server, c.port, typename)
	if len(values) > 0 {
		url += fmt.Sprintf("?%s", values.Encode())
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	content, ok := m[typename + "s"]
	if !ok {
		return nil, fmt.Errorf("No %ss in Response", typename)
	}
	var rlist []ListResult
	err = json.Unmarshal(*content, &rlist)
	return rlist, err
}

func (c *Client) List(typename string, count int) ([]ListResult, error) {
	return c.ListByParent(typename, "", 0)
}

func (c *Client) ListDetailByParent(
	typename string, parent_id string, fields []string, count int) (
		[]IObject, error) {
	var values url.Values
	values = make(url.Values, 0)
	if len(parent_id) > 0 {
		values.Add("parent_id", parent_id)
	}
	if len(fields) > 0 {
		values.Add("fields", strings.Join(fields, ","))
	}
	if count > 0 {
		values.Add("count", strconv.Itoa(count))
	}
	values.Add("detail", "true")

	url := fmt.Sprintf("http://%s:%d/%ss?%s",
		c.server, c.port, typename, values.Encode())
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	content, ok := m[typename + "s"]
	if !ok {
		return nil, fmt.Errorf("No %ss in Response", typename)
	}

	var elements []*json.RawMessage
	err = json.Unmarshal(*content, &elements)
	if err != nil {
		return nil, err
	}

	var result []IObject
	var xtype reflect.Type = typeMap[typename]

	for _, element := range elements {
		var item map[string]*json.RawMessage
		err = json.Unmarshal(*element, &item)
		if err != nil {
			return nil, err
		}

		content, ok := item[typename]
		if !ok {
			return nil, fmt.Errorf("No %s in element", typename)
		}

		valueT := reflect.New(xtype)
		obj := valueT.Interface().(IObject)
		err = json.Unmarshal(*content, obj)
		if err != nil {
			return nil, err
		}
		obj.SetClient(c)
		result = append(result, obj)
	}

	return result, nil
}

func (c *Client) ListDetail(typename string, fields []string, count int) (
	[]IObject, error) {
	return c.ListDetailByParent(typename, "", fields, count)
}

// Retrieve a specified field of an object from the API server.
func (c *Client) GetField(obj IObject, field string) error {
	url := fmt.Sprintf("%s?fields=%s", obj.GetHref(), field)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage
	err = json.Unmarshal(body, &m)

	if err != nil {
		return err
	}

	return json.Unmarshal(m[obj.GetType()], obj)
}

// Send a reference update message to the API server.
func (c *Client) UpdateReference(msg *ReferenceUpdateMsg) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s:%d/ref-update", c.server, c.port)
	resp, err := c.httpClient.Post(url, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}

func RegisterTypeMap(m TypeMap) {
	typeMap = m
}
