//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package config

import (
	"fmt"
	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	"strings"
)

func GetProjectId(
	client contrail.ApiClient, project_name string, project_id string) (
	string, error) {
	if len(project_id) > 0 {
		uuid := strings.ToLower(project_id)
		if !IsUuid(uuid) {
			return "",
				fmt.Errorf("Invalid uuid value: %s\n", uuid)
		}
		return uuid, nil
	}

	var name string
	if strings.ContainsRune(project_name, ':') {
		name = project_name
	} else {
		obj := new(types.Project)
		fqn := append(obj.GetDefaultParent(), project_name)
		name = strings.Join(fqn, `:`)
	}
	return client.UuidByName("project", name)
}

func GetProjectFQN(
	client contrail.ApiClient, projectName string, projectId string) (
	[]string, error) {
	if len(projectId) > 0 {
		uuid := strings.ToLower(projectId)
		if !IsUuid(uuid) {
			return nil,
				fmt.Errorf("Invalid uuid value: %s\n", uuid)
		}
		return client.FQNameByUuid(uuid)
	}

	if strings.ContainsRune(projectName, ':') {
		return strings.Split(projectName, ":"), nil
	}

	obj := new(types.Project)
	return append(obj.GetDefaultParent(), projectName), nil
}

// TODO: Create default security-group.
func CreateProject(client contrail.ApiClient, name string, createIpam bool) (
	string, error) {
	project := new(types.Project)
	project.SetName(name)
	err := client.Create(project)
	if err != nil {
		return "", err
	}
	if createIpam {
		ipam := new(types.NetworkIpam)
		ipam.SetParent(project)
		ipam.SetName("default-network-ipam")
		err = client.Create(ipam)
		if err != nil {
			client.Delete(project)
			return "", err
		}
	}
	return project.GetUuid(), nil
}

func DeleteProject(client contrail.ApiClient, project_id string) error {
	obj, err := client.FindByUuid("project", project_id)
	if err != nil {
		return err
	}
	defer client.Delete(obj)

	project := obj.(*types.Project)
	refList, err := project.GetNetworkIpams()
	for _, ref := range refList {
		client.DeleteByUuid("network-ipam", ref.Uuid)
	}
	return nil
}
