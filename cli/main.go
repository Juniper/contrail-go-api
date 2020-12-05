//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package main

import (
	"github.com/Juniper/contrail-go-api"
	"flag"
	"fmt"
	"os"
	"sort"
)

type ExecFunc func(client *contrail.Client, flagSet *flag.FlagSet)

type CliCommand struct {
	flagSet *flag.FlagSet
	exec ExecFunc
}

var (
	// OpenContrail API server
	oc_server string
	oc_port int

	// OpenContrail HTTPS control
	oc_insecure bool
	oc_skip_verify bool
	oc_ca_file string
	oc_key_file string
	oc_cert_file string

	// Authentication
	// os_auth_strategy string
	os_auth_url string
	os_tenant_name string
	os_tenant_id string
	os_username string
	os_password string
	os_token string

	// Authentication HTTPS control
	os_insecure bool
	os_skip_verify bool
	os_ca_file string
	os_key_file string
	os_cert_file string

	commandMap map[string]CliCommand = make(map[string]CliCommand, 0)
)

func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc) {
	commandMap[name] = CliCommand{flagSet, exec}
}

func InitFlags() {
	flag.StringVar(&oc_server, "server", "localhost", "OpenContrail API server hostname or address")
	flag.IntVar(&oc_port, "port", 8082, "OpenContrail API server port")

	flag.BoolVar(&oc_insecure, "insecure", false, "OpenContrail https control")
	flag.BoolVar(&oc_skip_verify, "skip-verify", false, "OpenContrail https skip verification control")
	flag.StringVar(&oc_ca_file, "ca-file", "", "OpenContrail https CA file")
	flag.StringVar(&oc_key_file, "key-file", "", "OpenContrail https key file")
	flag.StringVar(&oc_cert_file, "cert-file", "", "OpenContrail https cert file")

	// default_strategy := os.Getenv("OS_AUTH_STRATEGY")
	// if len(default_strategy) == 0 {
	// 	default_strategy = "keystone"
	// }
	// flag.StringVar(&os_auth_strategy,
	// 	"os-auth-strategy", default_strategy,
	// 	"Authentication strategy (Env: OS_AUTH_STRATEGY)")
	flag.StringVar(&os_auth_url, "os-auth-url", os.Getenv("OS_AUTH_URL"), "Authentication URL (Env: OS_AUTH_URL)")
	flag.StringVar(&os_tenant_name, "os-tenant-name", os.Getenv("OS_TENANT_NAME"), "Authentication tenant name (Env: OS_TENANT_NAME)")
	flag.StringVar(&os_tenant_id, "os-tenant-id", os.Getenv("OS_TENANT_ID"), "Authentication tenant id (Env: OS_TENANT_ID)")
	flag.StringVar(&os_username, "os-username", os.Getenv("OS_USERNAME"), "Authentication username (Env: OS_USERNAME)")
	flag.StringVar(&os_password, "os-password", os.Getenv("OS_PASSWORD"), "Authentication password (Env: OS_PASSWORD)")
	flag.StringVar(&os_token, "os-token", os.Getenv("OS_TOKEN"), "Authentication URL (Env: OS_TOKEN)")

	flag.BoolVar(&os_insecure, "os-insecure", false, "Authentication https control")
	flag.BoolVar(&os_skip_verify, "os-skip-verify", false, "Authentication https skip verification control")
	flag.StringVar(&os_ca_file, "os-ca-file", "", "Authentication https CA file")
	flag.StringVar(&os_key_file, "os-key-file", "", "Authentication https key file")
	flag.StringVar(&os_cert_file, "os-cert-file", "", "Authentication https cert file")
}

func setupAuthKeystone(client *contrail.Client) {
	keystone := contrail.NewKeystoneClient(
		os_auth_url,
		os_tenant_name,
		os_username,
		os_password,
		os_token,
	)
	if !os_insecure {
		keystone.AddEncryption(os_ca_file, os_key_file, os_cert_file, os_skip_verify)
	}
	err := keystone.Authenticate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client.SetAuthenticator(keystone)
}

func usage() {
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  Commands:\n")
	commandMapArray := make([]string, len(commandMap))
	i := 0
	for j, _ := range commandMap {
		commandMapArray[i] = j
		i++
	}
	sort.Strings(commandMapArray)
	for s, _ := range commandMapArray {
		fmt.Fprintf(os.Stderr, "    %s\n", commandMapArray[s])
	}
}

func main() {
	InitFlags()
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(2)
	}

	command := flag.Arg(0)
	cmd, inMap := commandMap[command]
	if !inMap {
		usage()
		os.Exit(2)
	}

	flagSet := cmd.flagSet
	flagSet.Parse(flag.Args()[1:])

	client := contrail.NewClient(oc_server, oc_port)
	if !oc_insecure {
		client.AddEncryption(oc_ca_file, oc_key_file, oc_cert_file, oc_skip_verify)
	}
	if len(os_auth_url) > 0 {
		setupAuthKeystone(client)
	}

	cmd.exec(client, flagSet)
}
