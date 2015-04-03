//
// Copyright (c) 2015 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"code.google.com/p/go-uuid/uuid"
	"github.com/Juniper/contrail-go-api"
)

type interfaceStatusOptions struct {
	vrouter string
	detail bool
}

var (
	interfaceStatusOpts interfaceStatusOptions
)

func interfaceStatus(client *contrail.Client, flagSet *flag.FlagSet) {
	if len(interfaceStatusOpts.vrouter) == 0 {
		fmt.Fprintln(os.Stderr, "virtual-router must be specified.")
		os.Exit(1)
	}
	url := fmt.Sprintf("http://%s:8085/Snh_ItfReq",
		interfaceStatusOpts.vrouter)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// see: controller/src/vnsw/agent/oper/agent.sandesh
	// struct ItfSandeshData
	type InterfaceData struct {
		Name string		`xml:"name"`
		Uuid string		`xml:"uuid"`
		VrfName string		`xml:"vrf_name"`
		Status string		`xml:"active"`
		NetworkName string	`xml:"vn_name"`
		InstanceName string	`xml:"vm_name"`
		IpAddress string	`xml:"ip_addr"`
		LinkLocalAddress string `xml:"mdata_ip_addr"`
	}

	type Envelope struct {
		XMLName xml.Name `xml:"ItfResp"`
		Data []InterfaceData `xml:"itf_list>list>ItfSandeshData"`
	}
	
	var m Envelope
	err = xml.Unmarshal(body, &m)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	const interfaceStatusTmpl = `
  Interface: {{.Uuid}}
    Instance: {{.InstanceName}}
    Network: {{.NetworkName}}
    IpAddress: {{.IpAddress}} Link Local: {{.LinkLocalAddress}}
    Status: {{.Status}}
`
	var writer *tabwriter.Writer
	var tmpl *template.Template

	if interfaceStatusOpts.detail {
		tmpl = template.Must(template.New("interface-status").Parse(
			interfaceStatusTmpl))
	} else {
		writer = new(tabwriter.Writer)
		writer.Init(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintf(writer, "Instance\tNetwork\tIpAddress\n")
	}
	for _, ifdata := range m.Data {
		id := uuid.Parse(ifdata.Uuid)
		if uuid.Equal(id, uuid.NIL) {
			continue
		}
		if interfaceStatusOpts.detail {
			tmpl.Execute(os.Stdout, ifdata)
		} else {
			netname := strings.Split(ifdata.NetworkName, ":")
			fmt.Fprintf(writer, "%.34s\t%-.28s\t%s\n",
				ifdata.InstanceName, netname[2], ifdata.IpAddress)
		}
	}

	if !interfaceStatusOpts.detail {
		writer.Flush()
	}
}

func init() {
	statusFlags := flag.NewFlagSet("interface-status", flag.ExitOnError)
	statusFlags.StringVar(&interfaceStatusOpts.vrouter, "virtual-router",
		"", "Virtual Router hostname or address")
	statusFlags.BoolVar(&interfaceStatusOpts.detail, "detail", false,
		"Detailed information")
	RegisterCliCommand("interface-status", statusFlags, interfaceStatus)
}
