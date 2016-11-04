//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package config

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	"github.com/golang/glog"
)

type NetworkInfo struct {
	Uuid         string
	Name         string
	AdminState   bool
	NetworkId    int
	Transit      bool
	Mode         string
	Subnets      []string
	Policies     []string
	RouteTargets []string
}

func buildNetworkInfo(net *types.VirtualNetwork, detail bool) (
	*NetworkInfo, error) {
	var subnets []string
	var policies []string

	refList, err := net.GetNetworkIpamRefs()
	if err != nil {
		return nil, err
	}

	for _, ref := range refList {
		attr := ref.Attr.(types.VnSubnetsType)
		for _, ipamSubnet := range attr.IpamSubnets {
			subnets = append(subnets, fmt.Sprintf("%s/%d",
				ipamSubnet.Subnet.IpPrefix,
				ipamSubnet.Subnet.IpPrefixLen))
		}
	}

	if detail {
		refList, err = net.GetNetworkPolicyRefs()
		for _, ref := range refList {
			policies = append(policies, strings.Join(ref.To, ":"))
		}
	}

	info := &NetworkInfo{
		net.GetUuid(),
		net.GetName(),
		net.GetIdPerms().Enable,
		net.GetVirtualNetworkProperties().NetworkId,
		net.GetVirtualNetworkProperties().AllowTransit,
		net.GetVirtualNetworkProperties().ForwardingMode,
		subnets,
		policies,
		net.GetRouteTargetList().RouteTarget,
	}
	return info, err
}

func NetworkShow(client contrail.ApiClient, uuid string, detail bool) (
	*NetworkInfo, error) {
	obj, err := client.FindByUuid("virtual-network", uuid)
	if err != nil {
		return nil, err
	}
	return buildNetworkInfo(obj.(*types.VirtualNetwork), detail)
}

func NetworkList(client contrail.ApiClient, project_id string, detail bool) (
	[]*NetworkInfo, error) {
	fields := []string{"network_ipams"}
	if detail {
		fields = append(fields, "network_policys")
	}
	networks, err := client.ListDetailByParent(
		"virtual-network", project_id, fields)
	if err != nil {
		return nil, err
	}

	var networkList []*NetworkInfo
	for _, reference := range networks {
		info, err := buildNetworkInfo(
			reference.(*types.VirtualNetwork), detail)
		if err != nil {
			return nil, err
		}
		networkList = append(networkList, info)
	}

	return networkList, nil
}

func makeSubnet(prefix string) (*types.IpamSubnetType, error) {
	expr := regexp.MustCompile(`(([0-9]{1,3}\.){3}[0-9]{1,3})/([0-9]{1,2})`)
	match := expr.FindStringSubmatch(prefix)
	if match == nil {
		return nil, fmt.Errorf("Invalid subnet prefix %s", prefix)
	}
	address := match[1]
	if net.ParseIP(address) == nil {
		return nil, fmt.Errorf("%s is not a valid IP address", address)
	}
	prefixlen, _ := strconv.Atoi(match[3])
	if prefixlen < 0 || prefixlen > 32 {
		return nil, fmt.Errorf("Invalid subnet prefix length %d", prefixlen)
	}

	subnet := &types.IpamSubnetType{
		Subnet: &types.SubnetType{address, prefixlen}}
	return subnet, nil
}

func networkAddSubnet(
	client contrail.ApiClient,
	project *types.Project, network *types.VirtualNetwork,
	subnet *types.IpamSubnetType,
	ipam *types.NetworkIpam) error {

	if ipam != nil {
		obj, err := client.FindByUuid("network-ipam", ipam.GetUuid())
		if err != nil {
			return err
		}
		ipam = obj.(*types.NetworkIpam)
	} else {
		obj, err := client.FindByName("network-ipam",
			"default-domain:default-project:default-network-ipam")
		if err != nil {
			return err
		}
		ipam = obj.(*types.NetworkIpam)
	}

	refs, err := network.GetNetworkIpamRefs()
	if err != nil {
		return err
	}

	var subnets types.VnSubnetsType

	for _, ref := range refs {
		if ref.Uuid == ipam.GetUuid() {
			subnets = ref.Attr.(types.VnSubnetsType)
			network.DeleteNetworkIpam(ref.Uuid)
			break
		}
	}
	subnets.AddIpamSubnets(subnet)
	network.AddNetworkIpam(ipam, subnets)
	return nil
}


func CreateNetworkWithSubnet(
client contrail.ApiClient, project_id, name, prefix string) (
string, error) {

	obj, err := client.FindByUuid("project", project_id)
	if err != nil {
		return "", err
	}

	project := obj.(*types.Project)

	net := new(types.VirtualNetwork)
	net.SetParent(project)
	net.SetName(name)

	subnet, err := makeSubnet(prefix)
	if err != nil {
		return "", err
	}
	err = networkAddSubnet(client, project, net, subnet, nil)
	if err != nil {
		return "", err
	}

	err = client.Create(net)
	if err != nil {
		return "", err
	}
	return net.GetUuid(), nil
}


func CreateNetworkWithIpam(
	client contrail.ApiClient, project *types.Project, name string, prefix[] string, ipams[] *types.NetworkIpam) (
	string, error) {

	net := new(types.VirtualNetwork)
	net.SetParent(project)
	net.SetName(name)

	for i, ipam := range ipams {
		subnet, err := makeSubnet(prefix[i])
		if err != nil {
			glog.Errorf("Cannot makeSubnet")
			continue
		}
		err = networkAddSubnet(client, project, net, subnet, ipam)
		if err != nil {
			glog.Errorf("Cannot add Ipam")
			continue
		}
	}

	err := client.Create(net)
	if err != nil {
		return "", err
	}
	return net.GetUuid(), nil
}

// AddSubnet
// returns true if the network was modified, false if the subnet already exists
// in the network.
func AddSubnet(
	client contrail.ApiClient, network *types.VirtualNetwork, prefix string) (
	bool, error) {

	ipamRefs, err := network.GetNetworkIpamRefs()
	if err != nil {
		return false, err
	}
	subnet, err := makeSubnet(prefix)
	if err != nil {
		return false, err
	}
	for _, ref := range ipamRefs {
		attr := ref.Attr.(types.VnSubnetsType)
		for _, entry := range attr.IpamSubnets {
			if *subnet.Subnet == *entry.Subnet {
				return false, nil
			}
		}
	}

	fqn := network.GetFQName()
	project, err := types.ProjectByName(client, strings.Join(fqn[:len(fqn)-1], ":"))
	if err != nil {
		return false, err
	}
	err = networkAddSubnet(client, project, network, subnet, nil)
	if err != nil {
		return false, err
	}
	err = client.Update(network)
	if err != nil {
		return false, err
	}
	return true, nil
}

func subnetTypeStringRepr(subnet *types.SubnetType) string {
	return fmt.Sprintf("%s/%d", subnet.IpPrefix, subnet.IpPrefixLen)
}

func RemoveSubnet(client contrail.ApiClient, network *types.VirtualNetwork, prefix string) error {
	ipamRefs, err := network.GetNetworkIpamRefs()
	if err != nil {
		return err
	}

	removeOp := func(ix int, uuid string, attr types.VnSubnetsType) error {
		ipam, err := types.NetworkIpamByUuid(client, uuid)
		if err != nil {
			return err
		}
		attr.IpamSubnets = append(attr.IpamSubnets[0:ix], attr.IpamSubnets[ix+1:]...)
		network.DeleteNetworkIpam(uuid)
		if len(attr.IpamSubnets) > 0 {
			network.AddNetworkIpam(ipam, attr)
		}
		return client.Update(network)
	}

	for _, ref := range ipamRefs {
		attr := ref.Attr.(types.VnSubnetsType)
		for ix, entry := range attr.IpamSubnets {
			entryPrefix := subnetTypeStringRepr(entry.Subnet)
			if entryPrefix == prefix {
				return removeOp(ix, ref.Uuid, attr)
			}
		}
	}

	return fmt.Errorf("Prefix %s not associated with network %s", prefix, network.GetName())
}

func CreateNetwork(client contrail.ApiClient, project_id, name string) (
	string, error) {

	obj, err := client.FindByUuid("project", project_id)
	if err != nil {
		return "", err
	}

	project := obj.(*types.Project)

	net := new(types.VirtualNetwork)
	net.SetParent(project)
	net.SetName(name)

	err = client.Create(net)
	if err != nil {
		return "", err
	}
	return net.GetUuid(), nil
}
