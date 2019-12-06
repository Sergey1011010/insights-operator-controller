package main

import (
	"github.com/redhatinsighs/insights-operator-controller/server"
	tests "github.com/redhatinsighs/insights-operator-controller/tests/rest"
	"github.com/verdverm/frisby"
)

func main() {
	if server.Environment != "production" {
		tests.ServerTests()
		tests.ClusterTests()
		tests.ConfigurationTests()
		tests.OperatorTests()
		tests.ProfileTests()
		tests.TriggerTests()
	} else {
		tests.AuthTests()
	}
	frisby.Global.PrintReport()
}
