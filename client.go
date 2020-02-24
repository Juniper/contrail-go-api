//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
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

// objectInterface defines the interface used internally between
// ObjectBase and Client implmementation
type objectInterface interface {
	GetField(IObject, string) error
	UpdateReference(*ReferenceUpdateMsg) error
}

// The Authenticator interface is used to add an autentication token on a per
// request basis. This is used by the Keystone authentication class to decorate
// the requests with a token.
type Authenticator interface {
	AddAuthentication(*http.Request) error
}

// NopAuthenticator is an authentication that doesn't modify the request.
type NopAuthenticator struct {
}

// AddAuthentication implements the Authenticator interface for NopAuthenticator.
func (*NopAuthenticator) AddAuthentication(*http.Request) error {
	return nil
}

// The Encryptor interface is used to add an encryption to the REST call
type Encryptor interface {
	AddEncryption(caFile string, keyFile string, certFile string, insecure bool) error
}

// NopEncryptor doesn't add encryption
type NopEncryptor struct {
}

// AddEncryption implements the Encryptor interface for NopEncryptor.
func (*NopEncryptor) AddEncryption(caFile string, keyFile string, certFile string, insecure bool) error {
	return nil
}

// AddEncryption implements the Encryptor interface for Client.
func (c *Client) AddEncryption(caFile string, keyFile string, certFile string, insecure bool) error {
	c.scheme = "https"
	tlsConfig := &tls.Config{}
	if insecure {
		tlsConfig.InsecureSkipVerify = true
	} else if caFile != "" {
		caCert, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
		if certFile != "" && keyFile != "" {
			cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			if err != nil {
				return nil
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	c.httpClient.Transport = transport

	return nil
}

// ApiClient interface
type ApiClient interface {
	Create(ptr IObject) error
	Update(ptr IObject) error
	DeleteByUuid(typename, uuid string) error
	Delete(ptr IObject) error
	FindByUuid(typename string, uuid string) (IObject, error)
	UuidByName(typename string, fqn string) (string, error)
	FQNameByUuid(uuid string) ([]string, error)
	FindByName(typename string, fqn string) (IObject, error)
	List(typename string) ([]ListResult, error)
	ListByParent(typename string, parentID string) ([]ListResult, error)
	ListDetail(typename string, fields []string) ([]IObject, error)
	ListDetailByParent(typename string, parentID string, fields []string) ([]IObject, error)
}

// APIServer list
type APIServer struct {
	server string
	next   *APIServer
}

// A Client of the OpenContrail API server.
type Client struct {
	servers    map[string]string
	scheme     string
	port       int
	httpClient *http.Client
	auth       Authenticator
	nextServer string
}

type TlsConfig struct {
	ca   string
	key  string
	cert string
}

// ListResult is the return type of the {List, ListByParent} API calls.
type ListResult struct {
	Fq_name []string
	Href    string
	Uuid    string
}

var (
	typeMap TypeMap
)

// NewClient allocates and initializes a Contrail API client.
//
func NewClient(server string, port int) *Client {
	client := new(Client)
	client.servers = map[string]string{server: server}
	client.port = port
	client.scheme = "http"
	client.httpClient = &http.Client{}
	client.auth = new(NopAuthenticator)
	client.nextServer = server
	return client
}

// NewClientPool initializes client with multiple API servers
func NewClientPool(servers []string, port int) *Client {
	client := NewClient(servers[0], port)
	for i, server := range servers {
		if i == len(servers)-1 {
			client.servers[server] = servers[0]
		} else {
			client.servers[server] = servers[i+1]
		}
	}

	rand.Seed(time.Now().UnixNano())
	client.nextServer = servers[(rand.Intn(len(servers)))]

	return client
}

// GetServer retrieves the name or address of the Contrail API server(round-robin).
func (c *Client) GetServer() string {
	server := c.nextServer
	c.nextServer = c.servers[server]
	return server
}

// SetAuthenticator enables the user to configure an Authenticator (e.g. Keystone)
// to be used by Contrail API requests.
func (c *Client) SetAuthenticator(auth Authenticator) {
	c.auth = auth
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

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	first, _, _ := net.SplitHostPort(req.URL.Host)
	for {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return resp, err
		}
		if resp.StatusCode == http.StatusBadGateway ||
			resp.StatusCode == http.StatusServiceUnavailable ||
			resp.StatusCode == http.StatusGatewayTimeout {
			host, _, _ := net.SplitHostPort(req.URL.Host)
			if c.servers[host] == first {
				return resp, err
			}
			req.URL.Host = fmt.Sprintf("%s:%d", c.servers[host], c.port)
		} else {
			return resp, err
		}
	}

}

func (c *Client) httpPost(url string, bodyType string, body io.Reader) (
	*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	err = c.auth.AddAuthentication(req)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) httpPut(url string, bodyType string, body io.Reader) (
	*http.Response, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	err = c.auth.AddAuthentication(req)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) httpGet(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	err = c.auth.AddAuthentication(req)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) httpDelete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	err = c.auth.AddAuthentication(req)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Create an object in the OpenContrail API server.
//
// The object must have been initialized with a name.
func (c *Client) Create(ptr IObject) error {
	xtype := typename(ptr)
	url := fmt.Sprintf("%s://%s:%d/%ss", c.scheme, c.GetServer(), c.port, xtype)

	objJson, err := json.Marshal(ptr)
	if err != nil {
		return err
	}

	var rawJson json.RawMessage = objJson
	msg := map[string]*json.RawMessage{
		xtype: &rawJson,
	}
	data, err := json.Marshal(msg)

	resp, err := c.httpPost(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	ptr.SetClient(c)

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
	resp, err := c.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
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
	msg := map[string]*json.RawMessage{
		ptr.GetType(): &rawJson,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := c.httpPut(ptr.GetHref(), "application/json",
		bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	err = ptr.UpdateReferences()
	if err != nil {
		return err
	}
	ptr.UpdateDone()

	return nil
}

// DeleteByUuid deletes the specified object.
func (c *Client) DeleteByUuid(typename, uuid string) error {
	url := fmt.Sprintf("%s://%s:%d/%s/%s", c.scheme, c.GetServer(), c.port, typename, uuid)
	resp, err := c.httpDelete(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	return nil
}

// Delete an object from the API server.
func (c *Client) Delete(ptr IObject) error {
	resp, err := c.httpDelete(ptr.GetHref())
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	return nil
}

// FindByUuid reads an object identified by UUID.
func (c *Client) FindByUuid(typename string, uuid string) (IObject, error) {
	url := fmt.Sprintf("%s://%s:%d/%s/%s", c.scheme, c.GetServer(), c.port,
		typename, uuid)
	return c.readObject(typename, url)
}

// UuidByName returns the UUID of an object as identified by its fully qualified name.
func (c *Client) UuidByName(typename string, fqn string) (string, error) {
	url := fmt.Sprintf("%s://%s:%d/fqname-to-id", c.scheme, c.GetServer(), c.port)
	request := struct {
		Typename string   `json:"type"`
		Fq_name  []string `json:"fq_name"`
	}{
		typename,
		strings.Split(fqn, ":"),
	}
	data, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	resp, err := c.httpPost(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s: %s", resp.Status, body)
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

// FQNameByUuid returns the fully-qualified name of an object as identified by a UUID.
func (c *Client) FQNameByUuid(uuid string) ([]string, error) {
	request := struct {
		Uuid string `json:"uuid"`
	}{
		uuid,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s://%s:%d/id-to-fqname", c.scheme, c.GetServer(), c.port)
	resp, err := c.httpPost(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var response struct {
		Type    string
		Fq_name []string
	}
	err = json.Unmarshal(body, &response)
	return response.Fq_name, err
}

// FindByName reads an object identified by fully-qualified name represented as a
// string.
func (c *Client) FindByName(typename string, fqn string) (IObject, error) {
	uuid, err := c.UuidByName(typename, fqn)
	if err != nil {
		return nil, err
	}
	href := fmt.Sprintf(
		"%s://%s:%d/%s/%s", c.scheme, c.GetServer(), c.port, typename, uuid)
	return c.readObject(typename, href)
}

// ListByParent retrieves the identifiers of the objects of a specific type that are
// descendents of a specific object.
func (c *Client) ListByParent(
	typename string, parentID string) ([]ListResult, error) {
	var values url.Values
	values = make(url.Values, 0)
	if len(parentID) > 0 {
		values.Add("parent_id", parentID)
	}

	url := fmt.Sprintf("%s://%s:%d/%ss", c.scheme, c.GetServer(), c.port, typename)
	if len(values) > 0 {
		url += fmt.Sprintf("?%s", values.Encode())
	}
	resp, err := c.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	content, ok := m[typename+"s"]
	if !ok {
		return nil, fmt.Errorf("No %ss in Response", typename)
	}
	var rlist []ListResult
	err = json.Unmarshal(*content, &rlist)
	return rlist, err
}

// List retrieves the identifiers of all objects of a given type.
func (c *Client) List(typename string) ([]ListResult, error) {
	return c.ListByParent(typename, "")
}

// ListDetailByParent reads all the objects of a given type that are descendents of the
// specified parent object.
func (c *Client) ListDetailByParent(
	typename string, parentID string, fields []string) (
	[]IObject, error) {
	var values url.Values
	values = make(url.Values, 0)
	if len(parentID) > 0 {
		values.Add("parent_id", parentID)
	}
	for _, field := range fields {
		values.Add("fields", field)
	}
	values.Add("detail", "true")

	url := fmt.Sprintf("%s://%s:%d/%ss?%s", c.scheme, c.GetServer(), c.port, typename, values.Encode())
	resp, err := c.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	content, ok := m[typename+"s"]
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

// ListDetail reads all the objects of a specific type.
func (c *Client) ListDetail(typename string, fields []string) (
	[]IObject, error) {
	return c.ListDetailByParent(typename, "", fields)
}

// GetField retrieves a specified field of an object from the API server.
// This API is used by the generated types library to retrieve reference lists.
func (c *Client) GetField(obj IObject, field string) error {
	url := fmt.Sprintf("%s?fields=%s", obj.GetHref(), field)
	resp, err := c.httpGet(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	var m map[string]json.RawMessage
	err = json.Unmarshal(body, &m)

	if err != nil {
		return err
	}

	return json.Unmarshal(m[obj.GetType()], obj)
}

// UpdateReference sends a reference update message to the API server.
// Used by the generated types library.
func (c *Client) UpdateReference(msg *ReferenceUpdateMsg) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s://%s:%d/ref-update", c.scheme, c.GetServer(), c.port)
	resp, err := c.httpPost(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	return nil
}

// RegisterTypeMap is used by the generated types library to register the list of known
// object types.
func RegisterTypeMap(m TypeMap) {
	typeMap = m
}
