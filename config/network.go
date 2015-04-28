//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package config

import (
	"fmt"
	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	"regexp"
	"strconv"
	"strings"
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
		"virtual-network", project_id, fields, 0)
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

func CreateNetworkWithSubnet(
	client contrail.ApiClient, project_id, name, prefix string) (
	string, error) {

	expr := regexp.MustCompile(`(([0-9]{1,3}\.){3}[0-9]{1,3})/([0-9]{1,2})`)
	match := expr.FindStringSubmatch(prefix)
	if match == nil {
		return "", fmt.Errorf("Invalid subnet prefix %s", prefix)
	}
	address := match[1]
	prefixlen, _ := strconv.Atoi(match[3])

	obj, err := client.FindByUuid("project", project_id)
	if err != nil {
		return "", err
	}

	project := obj.(*types.Project)
	refList, err := project.GetNetworkIpams()
	if err != nil {
		return "", err
	}

	var ipam *types.NetworkIpam
	if len(refList) > 0 {
		obj, err := client.FindByUuid("network-ipam", refList[0].Uuid)
		if err != nil {
			return "", err
		}
		ipam = obj.(*types.NetworkIpam)
	} else {
		obj, err := client.FindByName("network-ipam",
			"default-domain:default-project:default-network-ipam")
		if err != nil {
			return "", err
		}
		ipam = obj.(*types.NetworkIpam)
	}

	net := new(types.VirtualNetwork)
	net.SetParent(project)
	net.SetName(name)

	subnets := types.VnSubnetsType{}
	subnets.AddIpamSubnets(
		&types.IpamSubnetType{
			Subnet: &types.SubnetType{address, prefixlen}})
	net.AddNetworkIpam(ipam, subnets)
	err = client.Create(net)
	if err != nil {
		return "", err
	}
	return net.GetUuid(), nil
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
