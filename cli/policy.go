//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
	"github.com/Juniper/contrail-go-api/types"
	"github.com/pborman/uuid"
)

type policyCommonOptions struct {
	project   string
	projectId string
}

type policyOpOptions struct {
	policy string
}

type policyListOptions struct {
	allTenants bool
	detail     bool
}

type policyAttachOptions struct {
	policy  string
	network string
}

type policyProtocolValue string
type policyPortValue int

var protocolValues = map[string]int{
	"tcp":  6,
	"udp":  17,
	"icmp": 1,
	"any":  0,
}

type policyRuleOptions struct {
	ruleId       string
	afterRule    string
	policy       string
	srcIpAddress string
	srcNetwork   string
	dstIpAddress string
	dstNetwork   string
	protocol     policyProtocolValue
	srcPort      policyPortValue
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
	policyOpOpts     policyOpOptions
	policyListOpts   policyListOptions
	policyRuleOpts   policyRuleOptions
	policyAttachOpts policyAttachOptions
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

// Retrieves the virtual-network references from the policy rules
// for display purposes.
func getRulesNetworks(policy *types.NetworkPolicy) (string, string) {
	displayValue := func(m map[string]bool) string {
		if len(m) > 1 {
			return "<multiple>"
		}
		for key, _ := range m {
			fqn := strings.Split(key, ":")
			return fqn[len(fqn)-1]
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

// Summary information of the network policies configured (per project)
func policyListTerse(client *contrail.Client, projectId string) {
	poList, err := client.ListDetailByParent(
		"network-policy", projectId, nil)
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

func makeShowTemplate() *template.Template {
	fm := template.FuncMap{
		"join": strings.Join,
	}
	tmpl := template.Must(template.New("policy-list").Funcs(fm).Parse(
		policyShowTmpl))
	tmpl.Parse(policyShowDetail)
	return tmpl
}

// Display the rules associated with a policy as well as the attached networks.
func policyListDetail(client *contrail.Client, projectId string) {
	poList, err := client.ListDetailByParent(
		"network-policy", projectId,
		[]string{"virtual_network_back_refs"})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tmpl := makeShowTemplate()
	for _, policy := range poList {
		tmpl.Execute(os.Stdout, policy)
	}
}

// List all the policies under a specific project (or all projects if
// all-tenants is specified.
func policyList(client *contrail.Client, flagSet *flag.FlagSet) {
	var projectId string

	if !policyListOpts.allTenants {
		var err error
		projectId, err = config.GetProjectId(
			client, policyCommonOpts.project,
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

// Detailed output for a specific policy
func policyShow(client *contrail.Client, flagSet *flag.FlagSet) {
	policy, err := getPolicyObject(
		client,
		policyCommonOpts.project,
		policyCommonOpts.projectId,
		policyOpOpts.policy)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tmpl := makeShowTemplate()
	tmpl.Execute(os.Stdout, policy)
}

// Create (an empty) network-policy.
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

// Delete an existing network-policy
func policyDelete(client *contrail.Client, flagSet *flag.FlagSet) {
	if len(policyOpOpts.policy) == 0 {
		fmt.Fprintf(os.Stderr, "policy name or uuid must be specified")
		os.Exit(2)
	}

	uuid, err := getPolicyId(client, policyCommonOpts.project,
		policyCommonOpts.projectId, policyOpOpts.policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = client.DeleteByUuid("network-policy", uuid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Locate the policy given project name/id and policy name/id.
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

func getPolicyId(client *contrail.Client,
	projectName, projectId, policyNameOrId string) (
	string, error) {

	uuid := strings.ToLower(policyNameOrId)
	if !config.IsUuid(uuid) {
		fqn, err := config.GetProjectFQN(client, projectName, projectId)
		if err != nil {
			return "", err
		}
		fqn = append(fqn, policyNameOrId)
		uuid, err = client.UuidByName("network-policy",
			strings.Join(fqn, ":"))
		if err != nil {
			return "", err
		}
	}
	return uuid, nil
}

// Parse command line arguments and set the Address field in the rule.
func makeAddresses(optAddress, optNetwork string) []types.AddressType {
	addresses := make([]types.AddressType, 1)
	address := &addresses[0]

	if len(optAddress) > 0 {
		comp := strings.Split(optAddress, "/")
		if len(comp) != 2 {
			fmt.Fprintf(os.Stderr,
				"Expected IP prefix in the format n.n.n.n/n"+
					", got %s\n", optAddress)
			os.Exit(2)
		}
		matched, _ := regexp.MatchString(config.IpAddressPattern,
			comp[0])
		if !matched {
			fmt.Fprintf(os.Stderr, "Invalid IP address: %s\n",
				comp[0])
			os.Exit(2)
		}
		address.Subnet.IpPrefix = comp[0]
		var err error
		address.Subnet.IpPrefixLen, err = strconv.Atoi(comp[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid prefix length")
			os.Exit(2)
		}
	}
	if len(optNetwork) > 0 {
		address.VirtualNetwork = optNetwork
	} else {
		address.VirtualNetwork = "any"
	}
	return addresses
}

// Helper for ports.
// TODO: handle ranges.
func makePorts(optPort policyPortValue) []types.PortType {
	var portList []types.PortType
	if optPort > 0 {
		portList = make([]types.PortType, 1)
		portList[0] = types.PortType{
			int(optPort),
			int(optPort),
		}
	}
	return portList
}

func makePolicyRule(opts *policyRuleOptions) *types.PolicyRuleType {
	rule := new(types.PolicyRuleType)
	// RuleSequence
	rule.RuleUuid = uuid.NewRandom().String()
	rule.Direction = `<>`
	rule.Protocol = string(opts.protocol)
	rule.SrcAddresses = makeAddresses(opts.srcIpAddress, opts.srcNetwork)
	rule.DstAddresses = makeAddresses(opts.dstIpAddress, opts.dstNetwork)
	rule.SrcPorts = makePorts(opts.srcPort)
	rule.DstPorts = makePorts(opts.dstPort)
	if opts.actionDrop {
		rule.ActionList.SimpleAction = "drop"
	} else {
		rule.ActionList.SimpleAction = "pass"
	}
	return rule
}

// Add a rule to an existing policy
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

	insertRule := func(list []types.PolicyRuleType, i int,
		rule *types.PolicyRuleType) []types.PolicyRuleType {
		return append(list[:i+1],
			append([]types.PolicyRuleType{*rule},
				list[i+1:]...)...)
	}
	if uuid := policyRuleOpts.afterRule; len(uuid) > 0 {
		matched, _ := regexp.MatchString(config.UuidPattern, uuid)
		if !matched {
			fmt.Fprintf(os.Stderr, "Invalid rule uuid %s\n", uuid)
			os.Exit(2)
		}
		var ok bool
		for i, entry := range entries.PolicyRule {
			if entry.RuleUuid == uuid {
				entries.PolicyRule =
					insertRule(entries.PolicyRule, i, rule)
				ok = true
				break
			}
		}
		if !ok {
			fmt.Fprintf(os.Stderr, "Rule uuid %s not found\n", uuid)
			os.Exit(2)
		}
	} else {
		entries.AddPolicyRule(rule)
	}
	policy.SetNetworkPolicyEntries(&entries)
	err = client.Update(policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Update a rule.
func policyRuleUpdate(client *contrail.Client, flagSet *flag.FlagSet) {
	policy, err := getPolicyObject(
		client,
		policyCommonOpts.project,
		policyCommonOpts.projectId,
		policyRuleOpts.policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	uuid := policyRuleOpts.ruleId
	matched, _ := regexp.MatchString(config.UuidPattern, uuid)
	if !matched {
		if len(uuid) == 0 {
			fmt.Fprintln(os.Stderr, "Unspecified rule uuid")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid uuid: %s\n", uuid)
		}
		os.Exit(2)
	}

	rule := makePolicyRule(&policyRuleOpts)

	var ok bool
	entries := policy.GetNetworkPolicyEntries()
	for i, entry := range entries.PolicyRule {
		if entry.RuleUuid == uuid {
			entries.PolicyRule[i] = *rule
			ok = true
			break
		}
	}

	if !ok {
		fmt.Fprintf(os.Stderr, "Rule uuid %s not found\n", uuid)
		os.Exit(2)
	}

	policy.SetNetworkPolicyEntries(&entries)
	err = client.Update(policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

// Delete the specified policy rule
func policyRuleDelete(client *contrail.Client, flagSet *flag.FlagSet) {
	policy, err := getPolicyObject(
		client,
		policyCommonOpts.project,
		policyCommonOpts.projectId,
		policyRuleOpts.policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	uuid := policyRuleOpts.ruleId
	matched, _ := regexp.MatchString(config.UuidPattern, uuid)
	if !matched {
		if len(uuid) == 0 {
			fmt.Fprintln(os.Stderr, "Unspecified rule uuid")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid uuid: %s\n", uuid)
		}
		os.Exit(2)
	}

	deleteRule := func(list []types.PolicyRuleType, i int) []types.PolicyRuleType {
		return append(list[:i], list[i+1:]...)
	}

	var ok bool
	entries := policy.GetNetworkPolicyEntries()
	for i, entry := range entries.PolicyRule {
		if entry.RuleUuid == uuid {
			entries.PolicyRule =
				deleteRule(entries.PolicyRule, i)
			ok = true
			break
		}
	}

	if !ok {
		fmt.Fprintf(os.Stderr, "Rule uuid %s not found\n", uuid)
		os.Exit(2)
	}

	policy.SetNetworkPolicyEntries(&entries)
	err = client.Update(policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func policyAttachNetwork(client *contrail.Client) *types.VirtualNetwork {
	if len(policyAttachOpts.network) == 0 {
		fmt.Fprintf(os.Stderr, "network name or id must be specified")
		os.Exit(2)
	}
	uuid := policyAttachOpts.network
	if !config.IsUuid(uuid) {
		var err error
		uuid, err = getNetworkUuidByName(client,
			policyCommonOpts.project,
			policyCommonOpts.projectId,
			policyAttachOpts.network)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	obj, err := client.FindByUuid("virtual-network", uuid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return obj.(*types.VirtualNetwork)

}
func policyAttach(client *contrail.Client, flagSet *flag.FlagSet) {
	network := policyAttachNetwork(client)

	policy, err := getPolicyObject(client,
		policyCommonOpts.project,
		policyCommonOpts.projectId,
		policyAttachOpts.policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	network.AddNetworkPolicy(policy, types.VirtualNetworkPolicyType{})
	err = client.Update(network)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func policyDetach(client *contrail.Client, flagSet *flag.FlagSet) {
	network := policyAttachNetwork(client)
	policyId, err := getPolicyId(client,
		policyCommonOpts.project,
		policyCommonOpts.projectId,
		policyAttachOpts.policy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = network.DeleteNetworkPolicy(policyId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = client.Update(network)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
	policyInitCommonOptions(showFlags)
	showFlags.StringVar(&policyOpOpts.policy, "policy", "",
		"Policy name or uuid")
	RegisterCliCommand("policy-show", showFlags, policyShow)

	createFlags := flag.NewFlagSet("policy-create", flag.ExitOnError)
	policyInitCommonOptions(createFlags)
	createFlags.Usage = policyCreateUsage(createFlags)
	RegisterCliCommand("policy-create", createFlags, policyCreate)

	deleteFlags := flag.NewFlagSet("policy-delete", flag.ExitOnError)
	policyInitCommonOptions(deleteFlags)
	deleteFlags.StringVar(&policyOpOpts.policy, "policy", "",
		"Policy name or uuid")
	RegisterCliCommand("policy-delete", deleteFlags, policyDelete)

	ruleAddFlags := flag.NewFlagSet("policy-rule-add", flag.ExitOnError)
	policyInitCommonOptions(ruleAddFlags)
	policyRuleAddUpdateInitOptions(ruleAddFlags)
	ruleAddFlags.StringVar(&policyRuleOpts.afterRule, "after", "",
		"Add new rule after the rule with the specified uuid")
	RegisterCliCommand("policy-rule-add", ruleAddFlags, policyRuleAdd)

	ruleUpdateFlags := flag.NewFlagSet("policy-rule-update",
		flag.ExitOnError)
	policyInitCommonOptions(ruleUpdateFlags)
	policyRuleAddUpdateInitOptions(ruleUpdateFlags)
	ruleUpdateFlags.StringVar(&policyRuleOpts.ruleId, "rule", "",
		"Rule uuid to update")
	RegisterCliCommand("policy-rule-update", ruleUpdateFlags,
		policyRuleUpdate)

	ruleDeleteFlags := flag.NewFlagSet("policy-rule-delete",
		flag.ExitOnError)
	policyInitCommonOptions(ruleDeleteFlags)
	ruleDeleteFlags.StringVar(&policyRuleOpts.policy, "policy", "",
		"Policy name or uuid")
	ruleDeleteFlags.StringVar(&policyRuleOpts.ruleId, "rule", "",
		"Rule uuid to delete")
	RegisterCliCommand("policy-rule-delete", ruleDeleteFlags,
		policyRuleDelete)

	attachFlags := flag.NewFlagSet("policy-attach", flag.ExitOnError)
	policyInitCommonOptions(attachFlags)
	attachFlags.StringVar(&policyAttachOpts.policy, "policy", "",
		"Policy name or uuid")
	attachFlags.StringVar(&policyAttachOpts.network, "network", "",
		"Network name or uuid")
	RegisterCliCommand("policy-attach", attachFlags, policyAttach)

	detachFlags := flag.NewFlagSet("policy-detach", flag.ExitOnError)
	policyInitCommonOptions(detachFlags)
	detachFlags.StringVar(&policyAttachOpts.policy, "policy", "",
		"Policy name or uuid")
	detachFlags.StringVar(&policyAttachOpts.network, "network", "",
		"Network name or uuid")
	RegisterCliCommand("policy-detach", detachFlags, policyDetach)

}
