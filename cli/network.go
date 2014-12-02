//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"contrail-go-api"
	"contrail-go-api/config"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type CommonOptions struct {
	project string
	project_id string
}

type ShowOptions struct {
	allTenants bool
}

var (
	commonOpts CommonOptions
	showOpts ShowOptions
)


func NetworkShow(client *contrail.Client, flagSet *flag.FlagSet) {
	var parent_id string

	if !showOpts.allTenants {
		var err error
		parent_id, err =
			config.GetProjectId(
			client, commonOpts.project, commonOpts.project_id)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	networkList, err := config.NetworkList(client, parent_id)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(writer, "Network\tUuid\tSubnets")
	for _, n := range networkList {
		fmt.Fprintf(writer, "%s\t%s\t%s\n", n.Name, n.Uuid, n.Subnets)
	}
	writer.Flush()
}

func initCommonFlags(flagSet *flag.FlagSet) {
	defaultProject := os.Getenv("OS_TENANT_NAME")
	if len(defaultProject) == 0 {
		defaultProject = "admin"
	}
	
	flagSet.StringVar(&commonOpts.project, "project", defaultProject,
		"Project name (Env: OS_TENANT_NAME)")
	flagSet.StringVar(&commonOpts.project_id, "project-id",
		os.Getenv("OS_TENANT_ID"), "Project id (Env: OS_TENANT_ID)")
}

func init() {
	showFlags := flag.NewFlagSet("network-show", flag.ExitOnError)
	initCommonFlags(showFlags)
	showFlags.BoolVar(&showOpts.allTenants, "all-tenants", false,
		"Display networks for all tenants")

	RegisterCliCommand("network-show", showFlags, NetworkShow)
}
