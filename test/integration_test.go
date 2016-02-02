package contrail_test

import (
	"flag"
)

var (
	testApiServer string
	testApiPort   int
)

func runIntegrationTest() bool {
	return testApiServer != ""
}

func init() {
	flag.StringVar(&testApiServer, "contrail-api", "", "Contrail api server for integration tests")
	flag.IntVar(&testApiPort, "contrail-api-port", 8082, "Contrail api server port for integration tests")
}
