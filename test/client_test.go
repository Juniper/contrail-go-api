package contrail

import (
	"contrail-go-api"
	"contrail-go-api/types"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)
	elements, err := client.List("project")
	if err != nil {
		t.Fatal(err)
	}
	if len(elements) < 1 {
		t.Error("Empty project list")
	}

	for _, element := range elements {
		obj, err := client.ReadListResult("project", &element)
		if obj == nil {
			t.Fatal(err)
		}

		var project *types.Project = obj.(*types.Project)
		networks, err := project.GetVirtualNetworks()
		if len(networks) < 1 {
			t.Error("Empty virtual-network list")
		}

		for _, network := range networks {
			iObj, err := client.ReadReference(
				"virtual-network", &network)
			if iObj == nil {
				t.Fatal(err)
			}
			net := iObj.(*types.VirtualNetwork)
			fmt.Println(net.GetName())
		}
	}
}

func TestCreate(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)
	net := types.VirtualNetwork{}
	net.SetName("test")
	ipam, err := client.FindByName("network-ipam",
		"default-domain:default-project:default-network-ipam")
	if err != nil {
		t.Fatal(err)
	}
	subnets := types.VnSubnetsType{}
	subnets.AddIpamSubnets(
		&types.IpamSubnetType{
			Subnet: types.SubnetType{"10.0.0.0", 8}})
	net.AddNetworkIpam(ipam.(*types.NetworkIpam), subnets)
	err = client.Create(&net)
	if err != nil {
		t.Error(err)
	}
	if len(net.GetUuid()) == 0 {
		t.Error("No uuid assigned")
	}

	netPtr, err := client.FindByUuid("virtual-network", net.GetUuid())
	if err != nil {
		t.Error(err)
	} else {
		xnet := netPtr.(*types.VirtualNetwork)
		refs, err := xnet.GetNetworkIpamRefs()
		if err != nil {
			t.Error(err)
		} else if len(refs) > 0 {
			xnet.DeleteNetworkIpam(refs[0].Uuid)
			nsubnets := types.VnSubnetsType{}
			nsubnets.AddIpamSubnets(
				&types.IpamSubnetType{
					Subnet: types.SubnetType{
						"192.168.0.0", 16}})
			xnet.AddNetworkIpam(ipam.(*types.NetworkIpam), nsubnets)
			client.Update(xnet)
		} else {
			t.Error("Empty network-ipam reference list")
		}
	}

	err = client.Delete(&net)
	if err != nil {
		t.Error(err)
	}
}
