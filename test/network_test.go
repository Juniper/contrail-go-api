//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
	"github.com/Juniper/contrail-go-api/types"
)

func expectNetworkHasSubnet(t *testing.T, network *types.VirtualNetwork, prefix string) {
	refs, err := network.GetNetworkIpamRefs()
	if err == nil && len(refs) > 0 {
		attr := refs[0].Attr.(types.VnSubnetsType)
		if len(attr.IpamSubnets) > 0 {
			address := strings.Split(prefix, "/")[0]
			if attr.IpamSubnets[0].Subnet.IpPrefix != address {
				t.Errorf("expected %s, got %s", address, attr.IpamSubnets[0].Subnet.IpPrefix)
			}
		} else {
			t.Error("ipam reference has no subnets")
		}
	} else {
		t.Error("Network object doesn't have ipam references")
	}
}

func networkTestSetup(t *testing.T) (contrail.ApiClient, string) {
	client := newTestClient()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "test"})
	err := client.Create(project)
	if err != nil {
		t.Fatal(err)
	}

	return client, project.GetUuid()
}

func networkTestTeardown(client contrail.ApiClient) {
	network, err := types.VirtualNetworkByName(client, "default-domain:test:subnet-test")
	if err == nil {
		client.Delete(network)
	}
	project, err := types.ProjectByName(client, "default-domain:test")
	if err == nil {
		client.Delete(project)
	}
}

func TestAddSubnetInsert(t *testing.T) {
	client, projectId := networkTestSetup(t)

	defer networkTestTeardown(client)

	netId, err := config.CreateNetwork(client, projectId, "subnet-test")
	if err != nil {
		t.Fatal(err)
	}
	network, err := types.VirtualNetworkByUuid(client, netId)
	if err != nil {
		t.Fatal(err)
	}
	result, err := config.AddSubnet(client, network, "192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Error("AddSubnet returned false")
	}
	expectNetworkHasSubnet(t, network, "192.168.0.0/24")
}

func TestAddSubnetDuplicate(t *testing.T) {
	client, projectId := networkTestSetup(t)

	defer networkTestTeardown(client)

	netId, err := config.CreateNetworkWithSubnet(client, projectId, "subnet-test", "192.168.0.0/24")
	require.NoError(t, err)
	network, err := types.VirtualNetworkByUuid(client, netId)
	require.NoError(t, err)

	expectNetworkHasSubnet(t, network, "192.168.0.0/24")

	result, err := config.AddSubnet(client, network, "192.168.0.0/24")
	require.NoError(t, err)
	assert.False(t, result)
}

func TestAddSubnetSecond(t *testing.T) {
	client, projectId := networkTestSetup(t)
	defer networkTestTeardown(client)

	netId, err := config.CreateNetworkWithSubnet(client, projectId, "subnet-test", "192.168.0.0/24")
	require.NoError(t, err)
	network, err := types.VirtualNetworkByUuid(client, netId)
	require.NoError(t, err)

	expectNetworkHasSubnet(t, network, "192.168.0.0/24")

	result, err := config.AddSubnet(client, network, "192.168.1.0/24")
	require.NoError(t, err)
	assert.True(t, result)

	refs, err := network.GetNetworkIpamRefs()
	if err == nil && len(refs) > 0 {
		attr := refs[0].Attr.(types.VnSubnetsType)
		assert.Len(t, attr.IpamSubnets, 2)
	} else {
		t.Error("Network object doesn't have ipam references")
	}

}

func TestRemoveSubnetLastItem(t *testing.T) {
	client, projectId := networkTestSetup(t)
	defer networkTestTeardown(client)

	netId, err := config.CreateNetworkWithSubnet(client, projectId, "subnet-test", "192.168.0.0/24")
	require.NoError(t, err)
	network, err := types.VirtualNetworkByUuid(client, netId)
	require.NoError(t, err)

	err = config.RemoveSubnet(client, network, "192.168.0.0/24")
	assert.NoError(t, err)

	ipamRefs, err := network.GetNetworkIpamRefs()
	require.NoError(t, err)
	assert.Len(t, ipamRefs, 0)
}

func expectSubnetCount(t *testing.T, network *types.VirtualNetwork, count int) {
	ipamRefs, err := network.GetNetworkIpamRefs()
	require.NoError(t, err)
	assert.Len(t, ipamRefs, 1)
	attr := ipamRefs[0].Attr.(types.VnSubnetsType)
	assert.Len(t, attr.IpamSubnets, count)
	for i, subnet := range attr.IpamSubnets {
		fmt.Printf("subnet %2d %s/%d\n", i, subnet.Subnet.IpPrefix, subnet.Subnet.IpPrefixLen)
	}
}

func TestRemoveSubnetOne(t *testing.T) {
	client, projectId := networkTestSetup(t)
	defer networkTestTeardown(client)

	netId, err := config.CreateNetworkWithSubnet(client, projectId, "subnet-test", "192.168.0.0/24")
	require.NoError(t, err)
	network, err := types.VirtualNetworkByUuid(client, netId)
	require.NoError(t, err)

	_, err = config.AddSubnet(client, network, "192.168.1.0/24")
	require.NoError(t, err, "ADD second subnet")

	// This read is required because the API server modifies the subnet attribute inside
	// the ipam-refs and currently there is no way for us to map the cached data as invalid.
	network, err = types.VirtualNetworkByUuid(client, netId)
	require.NoError(t, err)
	expectSubnetCount(t, network, 2)

	err = config.RemoveSubnet(client, network, "192.168.0.0/24")
	assert.NoError(t, err, "Remove first subnet")

	// flush the cached information
	network, err = types.VirtualNetworkByUuid(client, netId)
	require.NoError(t, err)
	expectSubnetCount(t, network, 1)
}

func TestBadSubnetValues(t *testing.T) {
	client, projectId := networkTestSetup(t)

	defer func() {
		project, err := types.ProjectByName(client, "default-domain:test")
		if err == nil {
			client.Delete(project)
		}
	}()

	badValues := []string{
		"x/y",
		"256.0.0.0/8",
		"192.168.0.0/-1",
		"0.0.0.0/64",
	}
	for i, value := range badValues {
		_, err := config.CreateNetworkWithSubnet(client, projectId, fmt.Sprintf("subnet%2d", i), value)
		assert.Error(t, err, "%s", value)
	}
}

func TestNetworkList(t *testing.T) {
	client := newTestClient()

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

	netList, err := config.NetworkList(client, project_id, false)
	if err != nil {
		t.Fatal(err)
	}

	if len(netList) != len(data) {
		t.Errorf("Expected %d networks, got %d",
			len(data), len(netList))
	}

	for _, net := range netList {
		fmt.Printf("%s %s %s\n", net.Uuid, net.Name, net.Subnets)
		if data[net.Name].Subnet != net.Subnets[0] {
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
