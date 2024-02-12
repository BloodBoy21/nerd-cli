package main

import (
	"flag"
	"nerd-cli/helpers"
	"nerd-cli/services"
)

func main() {
	modules := make(map[string]*bool)
	modules["config"] = flag.Bool("config", false, "Module to start the config")
	modules["service"] = flag.Bool("service", false, "Module to start the service")
	flagsMap := helpers.GetOptionFlags()
	flag.Parse()
	module, err := services.GetModule(modules, flagsMap)
	if err != nil {
		panic(err)
	}
	module.GetOptions()
	module.Run()
}
