package main

import (
	"context"
	"flag"
	"log"

	"github.com/stuartc63/terraform-provider-simple1/simple1"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var version = "dev"

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), simple1.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/custom/simple1",
		Debug:   debug,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
