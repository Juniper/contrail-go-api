package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
    "text/template"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
)

type machineCommonOptions struct {
    project string
    project_id string
}

type machineListOptions struct {
	detail bool
}

type machineShowOptions struct {
	detail bool
}

var (
    machineCommonOpts machineCommonOptions
    machineListOpts machineListOptions
    machineShowOpts machineShowOptions
)

const machineShowFormat = `  Machine: {{.DisplayName}}
      Uuid: {{.Uuid}}
      Instance Ip: {{.InstanceIp}}
      Floating Ip: {{.FloatingIp}}
` 

func machineList(client *contrail.Client, flagSet *flag.FlagSet) {
    var parent_id string
	var writer *tabwriter.Writer
    var err error

    parent_id, err = config.GetProjectId(client, machineCommonOpts.project, machineCommonOpts.project_id)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

	machineList, err := config.MachineList(client, parent_id, machineListOpts.detail)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	writer = new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(writer, "Uuid\tInstance Ip\tFloating Ip\n")
	for _, n := range machineList {
		fmt.Fprintf(writer, "%s\t%s\t%s\n", n.Uuid, n.InstanceIp, n.FloatingIp)
	}
	writer.Flush()
}

func machineShow(client *contrail.Client, flagSet *flag.FlagSet) {
    var tmpl string
    var uuid string
    
    if flagSet.NArg() < 1 {
        fmt.Println("Usage: virtual-machine-show <instance uuid>")
        os.Exit(1)
    }
 
    nameOrId := flagSet.Args()[0]
    if config.IsUuid(nameOrId) {
        uuid = nameOrId
    } else {
        fmt.Println("Valid instance Uuid not provided") 
    }

    info, err := config.MachineShow(client, uuid, machineShowOpts.detail)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    
    tmpl = machineShowFormat
    t := template.Must(template.New("machine-show").Parse(tmpl))
    t.Execute(os.Stdout, info) 
}

func machineInitCommonFlags(flagSet *flag.FlagSet) {
    defaultProject := os.Getenv("OS_TENANT_NAME")
    if len(defaultProject) == 0 {
        defaultProject = "admin"
    }
    
    flagSet.StringVar(&machineCommonOpts.project, "project", defaultProject,
        "Project name (Env: OS_TENANT_NAME)")
    flagSet.StringVar(&machineCommonOpts.project_id, "project-id",
        os.Getenv("OS_TENANT_ID"), "Project id (Env: OS_TENANT_ID)")
}

func init() {
	listFlags := flag.NewFlagSet("virtual-machine-list", flag.ExitOnError)
	machineInitCommonFlags(listFlags)
	RegisterCliCommand("virtual-machine-list", listFlags, machineList)	

	showFlags := flag.NewFlagSet("virtual-machine-show", flag.ExitOnError)
	machineInitCommonFlags(showFlags)
	RegisterCliCommand("virtual-machine-show", showFlags, machineShow)
}

