//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package config

import (
	"contrail-go-api"
	"contrail-go-api/types"
	"fmt"
	"regexp"
	"strings"
)

const uuid_pattern = `([0-9a-z]{32})|([0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12})`
func GetProjectId(
	client *contrail.Client, project_name string, project_id string) (
		string, error) {
	if len(project_id) > 0 {
		uuid := strings.ToLower(project_id)
		matched, err := regexp.MatchString(uuid_pattern, uuid)
		if err != nil {
			return "", nil
		}
		if !matched {
			return "", fmt.Errorf("Invalid uuid value: %s\n", uuid)
		}
		return uuid, nil
	}

	var name string
	if strings.ContainsRune(project_name, ':') {
		name = project_name
	} else {
		obj := &types.Project{}
		fqn := append(obj.GetDefaultParent(), project_name)
		name = strings.Join(fqn, `:`)
	}
	return client.UuidByName("project", name)
}

// TODO: Create default security-group.
func CreateProject(client *contrail.Client, name string, createIpam bool) (
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

func DeleteProject(client *contrail.Client, project_id string) error {
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
