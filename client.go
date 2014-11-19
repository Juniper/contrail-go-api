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
	"reflect"
	"strings"
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
type client struct {
	server string
	port   int
	httpClient *http.Client
	typeMap TypeMap
}

// The Client List API returns an array of ListResult entries.
type ListResult struct {
	Fq_name []string
	Href string
	Uuid string
}

// Allocates and initialized a client.
//
// The typeMap parameter specifies a map of name, reflection Type values
// use to deserialize the data received from the server.
func NewClient(typeMap TypeMap, server string, port int) *client {
	client := new(client)
	client.server = server
	client.port = port
	client.httpClient = &http.Client{}
	client.typeMap = typeMap
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

func (c *client) GetHttpClient() *http.Client {
	return c.httpClient
}

// Create an object in the OpenContrail API server.
//
// The object must have been initialized with a name.
func (c *client) Create(ptr IObject) error {
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
func (c *client) readObject(typename string, href string) (IObject, error) {
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

	var m map[string]json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	var xtype reflect.Type = c.typeMap[typename]
	valueT := reflect.New(xtype)
	obj := valueT.Interface().(IObject)
	err = json.Unmarshal(m[typename], obj)
	if err != nil {
		return nil, err
	}
	obj.SetClient(c)
	return obj, err
}

// Given a ListResult, retrieve an object from the API server.
func (c *client) ReadListResult(
	typename string, result *ListResult) (IObject, error) {
	return c.readObject(typename, result.Href)
}

// Given a link reference, retrieve an object from the API server.
func (c *client) ReadReference(
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
func (c *client) Update(ptr IObject) error {
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
func (c *client) Delete(ptr IObject) error {
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
func (c *client) FindByUuid(typename string, uuid string) (IObject, error) {
	url := fmt.Sprintf("http://%s:%d/%s/%s", c.server, c.port,
		typename, uuid)
	return c.readObject(typename, url)
}

// Read an object identified by fully-qualified name represented as a
// string.
func (c *client) FindByName(typename string, fqn string) (IObject, error) {
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
		return nil, err
	}
 	resp, err := c.httpClient.Post(url, "application/json",
		bytes.NewReader(data))
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

	m := struct {
		Uuid string
	}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	href := fmt.Sprintf(
		"http://%s:%d/%s/%s", c.server, c.port, typename, m.Uuid)
	return c.readObject(typename, href)
}

// Retrieve the list of all elements of a specific type.
func (c *client) List(typename string) ([]ListResult, error) {
	url := fmt.Sprintf("http://%s:%d/%ss", c.server, c.port, typename)
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

	var m map[string]json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	var rlist []ListResult
	err = json.Unmarshal(m[typename + "s"], &rlist)
	return rlist, err
}

// Retrieve a specified field of an object from the API server.
func (c *client) GetField(obj IObject, field string) error {
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
func (c *client) UpdateReference(msg *ReferenceUpdateMsg) error {
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
