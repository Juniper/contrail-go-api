//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
	"github.com/Juniper/contrail-go-api/types"
)

type virtualRouterListOptions struct {
	detail bool
}

type virtualRouterCreateOptions struct {
	ipAddress string
}

var (
	virtualRouterListOpts virtualRouterListOptions
	virtualRouterCreateOpts virtualRouterCreateOptions
)

type virtualRouterInfo struct {
	Name string
	IpAddress string
}

const virtualRouterShowDetail = `  Name: {{.Name}}
    IpAddress: {{.IpAddress}}

`

func virtualRouterList(client *contrail.Client, flagSet *flag.FlagSet) {
	var fields []string
	detail := virtualRouterListOpts.detail
	routerList, err := client.ListDetail("virtual-router", fields, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	routerMap := make(map[string]*virtualRouterInfo, 0)

	for _, obj := range routerList {
		router := obj.(*types.VirtualRouter)
		routerMap[router.GetName()] = &virtualRouterInfo{
			router.GetName(),
			router.GetVirtualRouterIpAddress(),
		}
	}

	// TODO: Get analytics information

	var tmpl *template.Template
	var wr *tabwriter.Writer

	if detail {
		tmpl = template.Must(
			template.New("virtual-router-list").Parse(
				virtualRouterShowDetail))
	} else {
		wr = new(tabwriter.Writer)
		wr.Init(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintf(wr, "Hostname\tIpAddress\n")
	}

	for _, value := range routerMap {
		if detail {
			tmpl.Execute(os.Stdout, value)
		} else {
			fmt.Fprintf(wr, "%s\t%s\n", value.Name, value.IpAddress)
		}
	}

	if !detail {
		wr.Flush()
	}
}

func virtualRouterCreate(client *contrail.Client, flagSet *flag.FlagSet) {
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(2)
	}

	if len(virtualRouterCreateOpts.ipAddress) == 0 {
		fmt.Fprintf(os.Stderr,
			"The ip-address option must be specified")
		os.Exit(2)
	}
	matched, err := regexp.MatchString(config.IpAddressPattern,
		virtualRouterCreateOpts.ipAddress)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if !matched {
		fmt.Fprintf(os.Stderr, "Invalid IP address specified: %s\n",
			virtualRouterCreateOpts.ipAddress)
		os.Exit(2)
	}

	name := flagSet.Args()[0]
	vrouter := new(types.VirtualRouter)
	vrouter.SetName(name)
	vrouter.SetVirtualRouterIpAddress(virtualRouterCreateOpts.ipAddress)
	client.Create(vrouter)
}

func virtualRouterDelete(client *contrail.Client, flagSet *flag.FlagSet) {
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(2)
	}

	name := flagSet.Args()[0]

	obj := new(types.VirtualRouter)
	fqn := append(obj.GetDefaultParent(), name)
	uuid, err := client.UuidByName("virtual-router", strings.Join(fqn, ":"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = client.DeleteByUuid("virtual-router", uuid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func virtualRouterNameUsage(argument *flag.FlagSet) func() {
	flagSet := argument
	return func() {
		flagSet.PrintDefaults()
		fmt.Fprintf(os.Stderr, "    virtual-router-name\n")
	}
}

func init() {
	listFlags := flag.NewFlagSet("virtual-router-list", flag.ExitOnError)
	listFlags.BoolVar(&virtualRouterListOpts.detail, "detail", false,
		"Detailed information")
	RegisterCliCommand("virtual-router-list", listFlags, virtualRouterList)

	createFlags := flag.NewFlagSet("virtual-router-create",
		flag.ExitOnError)
	createFlags.StringVar(&virtualRouterCreateOpts.ipAddress,
		"ip-address", "", "IP address of the virtual-router")
	createFlags.Usage = virtualRouterNameUsage(createFlags)
	RegisterCliCommand("virtual-router-create", createFlags,
		virtualRouterCreate)

	deleteFlags := flag.NewFlagSet("virtual-router-delete",
		flag.ExitOnError)
	deleteFlags.Usage = virtualRouterNameUsage(createFlags)
	RegisterCliCommand("virtual-router-delete", deleteFlags,
		virtualRouterDelete)
}
