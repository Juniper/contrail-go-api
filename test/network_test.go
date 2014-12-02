//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail

import (
	"contrail-go-api"
	"contrail-go-api/config"
	"fmt"
	"testing"
)

func TestNetworkList(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)

	project_id, err := config.CreateProject(client, "demo", true)
	if err != nil {
		t.Fatal(err)
	}
	defer config.DeleteProject(client, project_id)

	type TestData struct {
		Subnet string
		InList bool
	}

	data := map[string]*TestData{
		"test1": &TestData{"192.168.0.0/24", false},
		"test2": &TestData{"192.168.1.0/24", false},
		"test3": &TestData{"10.0.0.0/8", false},
		"test4": &TestData{"100.64.0.0/24", false},
	}

	for name, value := range data {
		nid, err := config.CreateNetworkWithSubnet(
			client, project_id, name, value.Subnet)
		if err != nil {
			t.Fatal(err)
		}
		defer client.DeleteByUuid("virtual-network", nid)
	}

	netList, err := config.NetworkList(client, project_id)
	if err != nil {
		t.Fatal(err)
	}

	if len(netList) != len(data) {
		t.Errorf("Expected %d networks, got %d",
			len(data), len(netList))
	}

	for _, net := range netList {
		fmt.Printf("%s %s %s\n", net.Uuid, net.Name, net.Subnets)
		if data[net.Name].Subnet != net.Subnets {
			t.Errorf("%s subnet %s, expected %s", net.Name,
				net.Subnets, data[net.Name].Subnet)
		}
		data[net.Name].InList = true
	}
	for name, value := range data {
		if !value.InList {
			t.Errorf("%s not present", name)
		}
	}
}
