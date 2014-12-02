//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package config

import (
	"contrail-go-api"
	"contrail-go-api/types"
	"fmt"
	"regexp"
	"strconv"
)

type NetworkInfo struct {
	Uuid string
	Name string
	Subnets string
}

func NetworkList(client *contrail.Client, project_id string) (
	[]NetworkInfo, error) {
	networks, err := client.ListDetailByParent(
		"virtual-network", project_id, []string{"network_ipams"}, 0)
	if err != nil {
		return nil, err
	}

	var networkList []NetworkInfo
	for _, reference := range networks {
		net := reference.(*types.VirtualNetwork)
		var subnets string

		refList, err := net.GetNetworkIpamRefs()
		if err != nil {
			return nil, err
		}

		for _, ref := range refList {
			attr := ref.Attr.(types.VnSubnetsType)
			for _, ipamSubnet := range attr.IpamSubnets {
				if len(subnets) > 0 {
					subnets += ","
				}
				subnets +=  fmt.Sprintf("%s/%d",
					ipamSubnet.Subnet.IpPrefix,
					ipamSubnet.Subnet.IpPrefixLen)
			}
		}
		networkList = append(networkList, NetworkInfo{
			net.GetUuid(),
			net.GetName(),
			subnets,
		})
	}

	return networkList, nil
}

func CreateNetworkWithSubnet(
	client *contrail.Client, project_id, name, prefix string) (
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
			Subnet: types.SubnetType{address, prefixlen}})
	net.AddNetworkIpam(ipam, subnets)
	err = client.Create(net)
	if err != nil {
		return "", err
	}
	return net.GetUuid(), nil
}
