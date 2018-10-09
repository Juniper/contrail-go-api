//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
	"github.com/Juniper/contrail-go-api/types"
)

type networkCommonOptions struct {
	project string
	project_id string
}

type networkListOptions struct {
	allTenants bool
	brief bool
	detail bool
}

type networkCreateOptions struct {
	subnet string
}

type networkDeleteOptions struct {
	purge bool
}

type networkShowOptions struct {
	detail bool
}

var (
	networkCommonOpts networkCommonOptions
	networkListOpts networkListOptions
	networkCreateOpts networkCreateOptions
	networkDeleteOpts networkDeleteOptions
	networkShowOpts networkShowOptions
)

const networkShowBrief = `  Network: {{.Name}}
      Uuid: {{.Uuid}}        State: {{if .AdminState}}UP{{else}}DOWN{{end}}
      Subnets: {{range .Subnets}}{{.}} {{end}}
`
const networkShowDetail = `  Network: {{.Name}}
      Uuid: {{.Uuid}}        State: {{if .AdminState}}UP{{else}}DOWN{{end}}
      NetwordId: {{.NetworkId | printf "%-5d"}}    Mode: {{.Mode}}    Transit: {{.Transit}}
      Subnets: {{range .Subnets}}{{.}} {{end}}{{if .RouteTargets}}
      RouteTargets: {{range .RouteTargets}}{{.}} {{end}}{{end}}
      {{if .Policies}}Policies:{{end}}{{range .Policies}}
         {{.}}
      {{end}}
`

func networkList(client *contrail.Client, flagSet *flag.FlagSet) {
	var parent_id string

	if !networkListOpts.allTenants {
		var err error
		parent_id, err = config.GetProjectId(
			client,	networkCommonOpts.project,
			networkCommonOpts.project_id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	networkList, err := config.NetworkList(client, parent_id,
		networkListOpts.detail)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if networkListOpts.brief || networkListOpts.detail {
		var tmpl string
		if (networkListOpts.detail) {
			tmpl = networkShowDetail
		} else {
			tmpl = networkShowBrief
		}
		t := template.Must(template.New("network-list").Parse(tmpl))
		for _, n := range networkList {
			t.Execute(os.Stdout, n)
		}
	} else {
		// terse format (wide line)
		writer := new(tabwriter.Writer)
		writer.Init(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(writer, "Network\tUuid\tSubnets")
		for _, n := range networkList {
			fmt.Fprintf(writer, "%s\t%s\t%s\n", n.Name, n.Uuid,
				strings.Join(n.Subnets, ", "))
		}
		writer.Flush()
	}
}

func networkCreateUsage(argument *flag.FlagSet) func() {
	flagSet := argument
	return func() {
		flagSet.PrintDefaults()
		fmt.Fprintf(os.Stderr, "    network-name\n")
	}
}

func networkCreate(client *contrail.Client, flagSet *flag.FlagSet) {
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(2)
	}

	name := flagSet.Args()[0]

	parent_id, err := config.GetProjectId(
		client, networkCommonOpts.project, networkCommonOpts.project_id)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if len(networkCreateOpts.subnet) > 0 {
		config.CreateNetworkWithSubnet(client, parent_id, name,
			networkCreateOpts.subnet)
	} else {
		config.CreateNetwork(client, parent_id, name)
	}
}

func networkNameOrIdUsage(argument *flag.FlagSet) func() {
	flagSet := argument
	return func() {
		flagSet.PrintDefaults()
		fmt.Fprintf(os.Stderr, "    network-name or uuid\n")
	}
}

func getNetworkUuidByName(
	client *contrail.Client, project_name, project_id, name string) (
		string, error) {
	var fqn []string
	if len(project_id) > 0 {
		uuid := strings.ToLower(project_id)
		if !config.IsUuid(uuid) {
			fmt.Fprintf(os.Stderr,
				"Invalid project-id value: %s\n", uuid)
			os.Exit(2)
		}
		obj, err := client.FindByUuid("project", project_id)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}
		fqn = obj.GetFQName()
	} else {
		fqn = strings.Split(project_name, ":")
		if len(fqn) == 1 {
			obj := new(types.Project)
			fqn = append(obj.GetDefaultParent(), project_name)
		}
	}
	fqn = append(fqn, name)
	return client.UuidByName("virtual-network", strings.Join(fqn, ":"))
}

func networkDelete(client *contrail.Client, flagSet *flag.FlagSet) {
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(2)
	}
	nameOrId := flagSet.Args()[0]

	var uuid string
	if config.IsUuid(nameOrId) {
		uuid = nameOrId
	} else {
		var err error
		uuid, err = getNetworkUuidByName(client,
			networkCommonOpts.project,
			networkCommonOpts.project_id, nameOrId)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}
	// TODO: purge requires deleting all children before the object
	client.DeleteByUuid("virtual-network", uuid)
}

func networkShow(client *contrail.Client, flagSet *flag.FlagSet) {
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(2)
	}
	nameOrId := flagSet.Args()[0]

	var uuid string
	if config.IsUuid(nameOrId) {
		uuid = nameOrId
	} else {
		var err error
		uuid, err = getNetworkUuidByName(client,
			networkCommonOpts.project,
			networkCommonOpts.project_id, nameOrId)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	info, err := config.NetworkShow(client, uuid, networkShowOpts.detail)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	var tmpl string
	if (networkShowOpts.detail) {
		tmpl = networkShowDetail
	} else {
		tmpl = networkShowBrief
	}

	t := template.Must(template.New("network-show").Parse(tmpl))
	t.Execute(os.Stdout, info)
}

func initCommonFlags(flagSet *flag.FlagSet) {
	defaultProject := os.Getenv("OS_TENANT_NAME")
	if len(defaultProject) == 0 {
		defaultProject = "admin"
	}
	
	flagSet.StringVar(&networkCommonOpts.project, "project", defaultProject,
		"Project name (Env: OS_TENANT_NAME)")
	flagSet.StringVar(&networkCommonOpts.project_id, "project-id",
		os.Getenv("OS_TENANT_ID"), "Project id (Env: OS_TENANT_ID)")
}

func init() {
	listFlags := flag.NewFlagSet("network-list", flag.ExitOnError)
	initCommonFlags(listFlags)
	listFlags.BoolVar(&networkListOpts.allTenants, "all-tenants", false,
		"Display networks for all tenants")
	listFlags.BoolVar(&networkListOpts.brief, "brief", false,
		"Multiline format")
	listFlags.BoolVar(&networkListOpts.detail, "detail", false,
		"Multiline format (detailed information)")
	RegisterCliCommand("network-list", listFlags, networkList)

	createFlags := flag.NewFlagSet("network-create", flag.ExitOnError)
	initCommonFlags(createFlags)
	createFlags.StringVar(&networkCreateOpts.subnet, "subnet", "",
		"Subnet prefix for network")
	createFlags.Usage = networkCreateUsage(createFlags)
	RegisterCliCommand("network-create", createFlags, networkCreate)

	deleteFlags := flag.NewFlagSet("network-delete", flag.ExitOnError)
	initCommonFlags(deleteFlags)
	deleteFlags.BoolVar(&networkDeleteOpts.purge, "purge", false,
		"Delete all dependent objects")
	deleteFlags.Usage = networkNameOrIdUsage(deleteFlags)
	RegisterCliCommand("network-delete", deleteFlags, networkDelete)

	showFlags := flag.NewFlagSet("network-show", flag.ExitOnError)
	initCommonFlags(showFlags)
	showFlags.BoolVar(&networkShowOpts.detail, "detail", false,
		"Detail output")
	showFlags.Usage = networkNameOrIdUsage(showFlags)
	RegisterCliCommand("network-show", showFlags, networkShow)
}
