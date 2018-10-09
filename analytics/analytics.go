//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package analytics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

const AnalyticsDefaultPort = 8081

// defined in src/base/sandesh/process_info.sandesh
type ProcessStatus struct {
	ModuleId string	`json:"module_id"`
	State string
}

const AnalyticsVRouter = `analytics/uves/vrouter`

// defined in src/vnsw/agent/uve/vrouter.sandesh
type VRouterNodeStatus struct {
	ProcessStatus []ProcessStatus  `json:"process_status"`
}

type AnalyticsClient struct {
	server string
	port int
	httpClient *http.Client
}

func NewAnalyticsClient(server string, port int) *AnalyticsClient {
	client := new(AnalyticsClient)
	client.server = server
	client.port = port
	client.httpClient = new(http.Client)
	return client
}

func (client *AnalyticsClient) VirtualRouterList() ([]string, error) {
	type Reference struct {
		Href string
		Name string
	}
	url := fmt.Sprintf("http://%s:%d/%ss",
		client.server, client.port, AnalyticsVRouter)
	resp, err := client.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var referenceList []Reference
	err = json.Unmarshal(body, &referenceList)
	if err != nil {
		return nil, err
	}

	var vrouters []string
	for _, reference := range referenceList {
		vrouters = append(vrouters, reference.Name)
	}
	return vrouters, nil
}

func (client *AnalyticsClient) VirtualRouterStatus(name string) (
	string, error) {

	url := fmt.Sprintf("http://%s:%d/%s/%s?cfilt=NodeStatus",
		client.server, client.port, AnalyticsVRouter, name)
	resp, err := client.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		NodeStatus VRouterNodeStatus
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response.NodeStatus.ProcessStatus) > 0 {
		return response.NodeStatus.ProcessStatus[0].State, nil
	}
	return "null", nil
}
