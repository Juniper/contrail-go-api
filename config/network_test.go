package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/mocks"
	"github.com/Juniper/contrail-go-api/types"
)

func createTestClient(t *testing.T) (contrail.ApiClient, string) {
	client := new(mocks.ApiClient)
	client.Init()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "test"})
	err := client.Create(project)
	if err != nil {
		t.Fatal(err)
	}

	return client, project.GetUuid()
}

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

func TestAddSubnetInsert(t *testing.T) {
	client, projectId := createTestClient(t)

	netId, err := CreateNetwork(client, projectId, "subnet-test")
	if err != nil {
		t.Fatal(err)
	}
	network, err := types.VirtualNetworkByUuid(client, netId)
	if err != nil {
		t.Fatal(err)
	}
	result, err := AddSubnet(client, network, "192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Error("AddSubnet returned false")
	}
	expectNetworkHasSubnet(t, network, "192.168.0.0/24")
}

func TestAddSubnetDuplicate(t *testing.T) {
	client, projectId := createTestClient(t)
	netId, err := CreateNetworkWithSubnet(client, projectId, "subnet-test", "192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}
	network, err := types.VirtualNetworkByUuid(client, netId)
	if err != nil {
		t.Fatal(err)
	}

	expectNetworkHasSubnet(t, network, "192.168.0.0/24")

	result, err := AddSubnet(client, network, "192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}
	if result {
		t.Error("AddSubnet expected false, got true")
	}
}

func TestAddSubnetSecond(t *testing.T) {
	client, projectId := createTestClient(t)
	netId, err := CreateNetworkWithSubnet(client, projectId, "subnet-test", "192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}
	network, err := types.VirtualNetworkByUuid(client, netId)
	if err != nil {
		t.Fatal(err)
	}

	expectNetworkHasSubnet(t, network, "192.168.0.0/24")

	result, err := AddSubnet(client, network, "192.168.1.0/24")
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Error("AddSubnet expected true, got false")
	}

	refs, err := network.GetNetworkIpamRefs()
	if err == nil && len(refs) > 0 {
		attr := refs[0].Attr.(types.VnSubnetsType)
		if len(attr.IpamSubnets) != 2 {
			t.Errorf("expected 2 subnets, got %d", len(attr.IpamSubnets))
		}
	} else {
		t.Error("Network object doesn't have ipam references")
	}

}

func TestBadSubnetValues(t *testing.T) {
	client, projectId := createTestClient(t)
	badValues := []string{
		"x/y",
		"256.0.0.0/8",
		"192.168.0.0/-1",
		"0.0.0.0/64",
	}
	for i, value := range badValues {
		_, err := CreateNetworkWithSubnet(client, projectId, fmt.Sprintf("subnet%2d", i), value)
		if err == nil {
			t.Errorf("expected error for %s", value)
		}
	}
}
