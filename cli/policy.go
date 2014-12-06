//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
	"github.com/Juniper/contrail-go-api/types"
)

type policyCommonOptions struct {
	project string
	projectId string
}

type policyListOptions struct {
	allTenants bool
	detail bool
}

type policyProtocolValue string
type policyPortValue int

var protocolValues = map[string]int{
	"tcp": 6,
	"udp": 17,
	"icmp": 1,
	"any": 0,
}

type policyRuleOptions struct {
	policy string
	srcIpAddress string
	srcNetwork string
	dstIpAddress string
	dstNetwork string
	protocol policyProtocolValue
	srcPort policyPortValue
	// srcPortRange string
	dstPort policyPortValue
	// dstPortRange string
	actionDrop bool
	//actionDeny bool
	//actionLog bool
	actionPass bool
	//actionReject bool
}

var (
	policyCommonOpts policyCommonOptions
	policyListOpts policyListOptions
	policyRuleOpts policyRuleOptions
)

const policyShowTmpl = `
{{define "AddressList"}}{{range .}}{{if .Subnet.IpPrefixLen}}{{.Subnet.IpPrefix}}/{{.Subnet.IpPrefixLen}}{{end}}{{.VirtualNetwork}}{{.SecurityGroup}}{{.NetworkPolicy}}{{end}}{{end}}
{{define "PortList"}}{{range .}}[{{.StartPort}}, {{.EndPort}}]{{end}}{{end}}
{{define "ActionList"}}{{.SimpleAction}}{{if .MirrorTo.AnalyzerName}}Analyzer({{.MirrorTo.AnalyzerName}}){{end}}{{end}}
`
const policyShowDetail = `  Policy: {{.GetName}}
    Uuid: {{.GetUuid}}{{range .GetNetworkPolicyEntries.PolicyRule}}
        Rule: {{.RuleUuid}}
            Source: {{template "AddressList" .SrcAddresses}}
            Destination: {{template "AddressList" .DstAddresses}}
            Protocol: {{.Protocol}}
            Source Ports: {{template "PortList" .SrcPorts}}
            Destination Ports: {{template "PortList" .DstPorts}}
	    Action: {{template "ActionList" .ActionList}}{{end}}

    Attached Networks:{{range .GetVirtualNetworkBackRefs}}
        {{join .To ":"}}{{end}}

`

func getRulesNetworks(policy *types.NetworkPolicy) (string, string) {
	displayValue := func (m map[string]bool) string {
		if len(m) > 1 {
			return "<multiple>"
		}
		for key, _ := range m {
			fqn := strings.Split(key, ":")
			return fqn[len(fqn) - 1]
		}
		return "none"
	}

	sourceMap := make(map[string]bool, 0)
	destMap := make(map[string]bool, 0)
	for _, rule := range policy.GetNetworkPolicyEntries().PolicyRule {
		if len(rule.SrcAddresses) > 0 &&
			len(rule.SrcAddresses[0].VirtualNetwork) > 0 {
			sourceMap[rule.SrcAddresses[0].VirtualNetwork] = true
		}
		if len(rule.DstAddresses) > 0 &&
			len(rule.DstAddresses[0].VirtualNetwork) > 0 {
			destMap[rule.DstAddresses[0].VirtualNetwork] = true
		}
	}

	source := displayValue(sourceMap)
	destination := displayValue(destMap)
	return source, destination
}

func policyListTerse(client *contrail.Client, projectId string) {
	poList, err := client.ListDetailByParent(
		"network-policy", projectId, nil, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(writer, "Uuid\tPolicy\tSource\tDestination")
	for _, obj := range poList {
		policy := obj.(*types.NetworkPolicy)
		source, destination := getRulesNetworks(policy)
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n",
			policy.GetUuid(), policy.GetName(), source, destination)
	}
	writer.Flush()
}

func policyListDetail(client *contrail.Client, projectId string) {
	poList, err := client.ListDetailByParent(
		"network-policy", projectId,
		[]string{"virtual_network_back_refs"}, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fm := template.FuncMap{
		"join": strings.Join,
	}
	tmpl := template.Must(template.New("policy-list").Funcs(fm).Parse(
		policyShowTmpl))
	tmpl.Parse(policyShowDetail)
	for _, policy := range poList {
		tmpl.Execute(os.Stdout, policy)
	}
}

func policyList(client *contrail.Client, flagSet *flag.FlagSet) {
	var projectId string

	if !policyListOpts.allTenants {
		var err error
		projectId, err = config.GetProjectId(
			client,	policyCommonOpts.project,
			policyCommonOpts.projectId)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if policyListOpts.detail {
		policyListDetail(client, projectId)
	} else {
		policyListTerse(client, projectId)
	}
}

func policyShow(client *contrail.Client, flagSet *flag.FlagSet) {
}

func policyCreate(client *contrail.Client, flagSet *flag.FlagSet) {
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(2)
	}

	name := flagSet.Args()[0]

	projectFQN, err := config.GetProjectFQN(client,
		policyCommonOpts.project, policyCommonOpts.projectId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	policy := new(types.NetworkPolicy)
	policy.SetFQName("project", append(projectFQN, name))
	err = client.Create(policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func policyDelete(client *contrail.Client, flagSet *flag.FlagSet) {
}

func getPolicyObject(
	client *contrail.Client,
	projectName, projectId, policyNameOrId string) (
	*types.NetworkPolicy, error) {

	if len(policyNameOrId) == 0 {
		return nil, fmt.Errorf("policy name or uuid must be specified")
	}

	if uuid := strings.ToLower(policyNameOrId); config.IsUuid(uuid) {
		obj, err := client.FindByUuid("network-policy", uuid)
		if err != nil {
			return nil, err
		}
		return obj.(*types.NetworkPolicy), nil
	}

	fqn, err := config.GetProjectFQN(client, projectName, projectId)
	if err != nil {
		return nil, err
	}
	fqn = append(fqn, policyNameOrId)
	obj, err := client.FindByName("network-policy", strings.Join(fqn, ":"))
	if err != nil {
		return nil, err
	}
	return obj.(*types.NetworkPolicy), nil
}

func makePolicyRule(opts *policyRuleOptions) *types.PolicyRuleType {
	rule := new(types.PolicyRuleType)
	// RuleSequence
	// RuleUuid
	rule.Direction = `<>`
	rule.Protocol = fmt.Sprintf("%d", opts.protocol)
	if len(opts.srcIpAddress) > 0 {
	} else if len(opts.srcNetwork) > 0 {
	}
	if opts.srcPort > 0 {
		rule.SrcPorts = make([]types.PortType, 1)
		rule.SrcPorts[0] = types.PortType{
			int(opts.srcPort),
			int(opts.srcPort),
		}
	}
	if opts.actionDrop {
		rule.ActionList.SimpleAction = "drop"
	} else {
		rule.ActionList.SimpleAction = "pass"
	}
	return rule
}

func policyRuleAdd(client *contrail.Client, flagSet *flag.FlagSet) {
	policy, err := getPolicyObject(
		client,
		policyCommonOpts.project,
		policyCommonOpts.projectId,
		policyRuleOpts.policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	entries := policy.GetNetworkPolicyEntries()
	rule := makePolicyRule(&policyRuleOpts)

	entries.AddPolicyRule(rule)
	policy.SetNetworkPolicyEntries(&entries)
	err = client.Update(policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func policyRuleUpdate(client *contrail.Client, flagSet *flag.FlagSet) {
}

func policyRuleDelete(client *contrail.Client, flagSet *flag.FlagSet) {
}

func policyInitCommonOptions(flagSet *flag.FlagSet) {
	defaultProject := os.Getenv("OS_TENANT_NAME")
	if len(defaultProject) == 0 {
		defaultProject = "admin"
	}

	flagSet.StringVar(&policyCommonOpts.project, "project", defaultProject,
		"Project name (Env: OS_TENANT_NAME)")
	flagSet.StringVar(&policyCommonOpts.projectId, "project-id",
		os.Getenv("OS_TENANT_ID"), "Project id (Env: OS_TENANT_ID)")
}

func policyRuleAddUpdateInitOptions(flagSet *flag.FlagSet) {
	flagSet.StringVar(&policyRuleOpts.policy, "policy", "",
		"Policy name or uuid")
	flagSet.StringVar(&policyRuleOpts.srcIpAddress, "source-address", "",
		"Source IP address prefix")
	flagSet.StringVar(&policyRuleOpts.dstIpAddress, "destination-address",
		"", "Destination IP address prefix")
	flagSet.StringVar(&policyRuleOpts.srcNetwork, "source-network", "",
		"Virtual-network name for source IP address")
	flagSet.StringVar(&policyRuleOpts.dstNetwork, "destination-network", "",
		"Virtual-network name for destination IP address")
	var allowedProtocolValues []string
	for key, _ := range protocolValues {
		allowedProtocolValues = append(allowedProtocolValues, key)
	}
	flagSet.Var(&policyRuleOpts.protocol, "protocol",
		fmt.Sprintf("IP protocol (one of [%s])",
			strings.Join(allowedProtocolValues, ",")))
	flagSet.Var(&policyRuleOpts.srcPort, "source-port",
		"Transport protocol (e.g. tcp, udp) source port")
	flagSet.Var(&policyRuleOpts.dstPort, "destination-port",
		"Transport protocol (e.g. tcp, udp) destination port")
	flagSet.BoolVar(&policyRuleOpts.actionDrop, "drop", false,
		"Discard packets")
	flagSet.BoolVar(&policyRuleOpts.actionDrop, "pass", false,
		"Accept packets (default)")
}

func policyCreateUsage(argument *flag.FlagSet) func() {
	flagSet := argument
	return func() {
		flagSet.PrintDefaults()
		fmt.Fprintf(os.Stderr, "    policy-name\n")
	}
}

func (p *policyProtocolValue) Set(value string) error {
	_, ok := protocolValues[value]
	if !ok {
		return fmt.Errorf("Invalid protocol value: %s", value)
	}
	*p = policyProtocolValue(value)
	return nil
}

func (p *policyProtocolValue) String() string {
	return string(*p)
}

func (p *policyProtocolValue) Get() interface{} {
	return string(*p)
}

func (p *policyPortValue) Set(value string) error {
	port, err := strconv.Atoi(value)
	if err != nil {
		port, err = net.LookupPort("", value)
		if err != nil {
			return err
		}
	}
	*p = policyPortValue(port)
	return nil
}

func (p *policyPortValue) String() string {
	return fmt.Sprintf("%d", int(*p))
}

func (p *policyPortValue) Get() interface{} {
	return int(*p)
}

func init() {
	listFlags := flag.NewFlagSet("policy-list", flag.ExitOnError)
	policyInitCommonOptions(listFlags)
	listFlags.BoolVar(&policyListOpts.allTenants,
		"all-tenants", false, "Display policies for all tenants")
	listFlags.BoolVar(&policyListOpts.detail,
		"detail", false, "Detailed information")
	RegisterCliCommand("policy-list", listFlags, policyList)

	showFlags := flag.NewFlagSet("policy-show", flag.ExitOnError)
	RegisterCliCommand("policy-show", showFlags, policyShow)

	createFlags := flag.NewFlagSet("policy-create", flag.ExitOnError)
	policyInitCommonOptions(createFlags)
	createFlags.Usage = policyCreateUsage(createFlags)
	RegisterCliCommand("policy-create", createFlags, policyCreate)

	deleteFlags := flag.NewFlagSet("policy-delete", flag.ExitOnError)
	RegisterCliCommand("policy-delete", deleteFlags, policyDelete)

	ruleAddFlags := flag.NewFlagSet("policy-rule-add", flag.ExitOnError)
	policyInitCommonOptions(ruleAddFlags)
	policyRuleAddUpdateInitOptions(ruleAddFlags)
	RegisterCliCommand("policy-rule-add", ruleAddFlags, policyRuleAdd)

	ruleUpdateFlags := flag.NewFlagSet("policy-rule-update",
		flag.ExitOnError)
	policyRuleAddUpdateInitOptions(ruleUpdateFlags)
	RegisterCliCommand("policy-rule-update", ruleUpdateFlags,
		policyRuleUpdate)

	ruleDeleteFlags := flag.NewFlagSet("policy-rule-delete",
		flag.ExitOnError)
	RegisterCliCommand("policy-rule-delete", ruleDeleteFlags,
		policyRuleDelete)
}
