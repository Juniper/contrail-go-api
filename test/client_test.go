package contrail

import (
	"contrail-go-api"
	"contrail-go-api/types"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)
	elements, err := client.List("project", 0)
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
			xattr := refs[0].Attr.(types.VnSubnetsType)
			if len(xattr.IpamSubnets) == 1 {
				isubnet := xattr.IpamSubnets[0]
				if isubnet.Subnet.IpPrefix != "10.0.0.0" ||
					isubnet.Subnet.IpPrefixLen != 8 {
					t.Errorf("Bad subnet %s/%d",
						isubnet.Subnet.IpPrefix,
						isubnet.Subnet.IpPrefixLen)
				}
			} else {
				t.Errorf("%d subnets", len(xattr.IpamSubnets))
			}
		} else {
			t.Error("Empty network-ipam reference list")
		}
	}

	err = client.Delete(&net)
	if err != nil {
		t.Error(err)
	}
}

func TestPropertyUpdate(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)
	net := types.VirtualNetwork{}
	net.SetName("test")

	err := client.Create(&net)
	if err != nil {
		t.Error(err)
	}

	defer client.Delete(&net)

	props := net.GetVirtualNetworkProperties()
	if len(props.ForwardingMode) > 0 {
		t.Error(props.ForwardingMode)
	}
	props.ForwardingMode = "l2_l3"
	net.SetVirtualNetworkProperties(&props)
	err = client.Update(&net)
	if err != nil {
		t.Error(err)
	}

	obj, err := client.FindByUuid("virtual-network", net.GetUuid())
	if err != nil {
		t.Fatal(err)
	}
	netPtr := obj.(*types.VirtualNetwork)
	p2 := netPtr.GetVirtualNetworkProperties()
	if p2.ForwardingMode != "l2_l3" {
		t.Errorf("Expected: l2_l3 got: %s", p2.ForwardingMode)
	}
}

func TestReferenceUpdate(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)
	net := types.VirtualNetwork{}
	net.SetName("test2")
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

	defer client.Delete(&net)

	obj, err := client.FindByUuid("virtual-network", net.GetUuid())
	if err != nil {
		t.Fatal(err)
	}

	netPtr := obj.(*types.VirtualNetwork)
	refs, err := netPtr.GetNetworkIpamRefs()
	if err != nil {
		t.Error(err)
	} else if len(refs) > 0 {
		netPtr.DeleteNetworkIpam(refs[0].Uuid)
		nsubnets := types.VnSubnetsType{}
		nsubnets.AddIpamSubnets(
			&types.IpamSubnetType{
				Subnet: types.SubnetType{
					"192.168.0.0", 16}})
		netPtr.AddNetworkIpam(ipam.(*types.NetworkIpam), nsubnets)
		client.Update(netPtr)
	} else {
		t.Error("Empty network-ipam reference list")
	}

	obj, err = client.FindByUuid("virtual-network", net.GetUuid())
	if err != nil {
		t.Fatal(err)
	}
	netPtr = obj.(*types.VirtualNetwork)
	refs, err = netPtr.GetNetworkIpamRefs()
	if err != nil {
		t.Error(err)
	} else if len(refs) != 1 {
		t.Errorf("Expected: 1 reference, %d actual", len(refs))
	} else {
		xattr := refs[0].Attr.(types.VnSubnetsType)
		if len(xattr.IpamSubnets) == 1 {
			isubnet := xattr.IpamSubnets[0]
			if isubnet.Subnet.IpPrefix != "192.168.0.0" ||
				isubnet.Subnet.IpPrefixLen != 16 {
				t.Errorf("Bad subnet %s/%d",
					isubnet.Subnet.IpPrefix,
					isubnet.Subnet.IpPrefixLen)
			}
		} else {
			t.Errorf("%d subnets", len(xattr.IpamSubnets))
		}
	}

	netPtr.ClearNetworkIpam()
	client.Update(netPtr)

	obj, err = client.FindByUuid("virtual-network", net.GetUuid())
	if err != nil {
		t.Fatal(err)
	}
	netPtr = obj.(*types.VirtualNetwork)
	refs, err = netPtr.GetNetworkIpamRefs()
	if err != nil {
		t.Error(err)
	} else if len(refs) != 0 {
		t.Errorf("Expected: 1 reference, %d actual", len(refs))
	}
}

func TestListDetail(t *testing.T) {
	client := contrail.NewClient("localhost", 8082)
	objects, err := client.ListDetail("virtual-network", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(objects) < 3 {
		t.Error("Default networks not present")
	}

	objectMap := make(map[string]*types.VirtualNetwork, 0)
	for _, object := range objects {
		net := object.(*types.VirtualNetwork)
		objectMap[net.GetName()] = net
	}

	expected := []string {
		"ip-fabric", "__link_local__",
	}
	for _, expect := range expected {
		_, ok := objectMap[expect]
		if !ok {
			t.Errorf("%s not found", expect)
		}
	}
}
