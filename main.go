package main

import (
	"ewallet-framework/cmd"
	"ewallet-framework/helpers"
)

func main() {
	// setup config
	helpers.SetupConfig()

	// setup log
	helpers.SetupLogger()

	// load database
	// helpers.SetupMySql()
	// run grpc
	go cmd.ServeGRPC()
	// run http
	cmd.ServeHTTP()
}
