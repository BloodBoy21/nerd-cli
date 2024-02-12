package main

import (
	"flag"
	"fmt"
	"nerd-cli/helpers"
	"nerd-cli/services"
	"os"
)

func initEmpty(){
	if len(os.Args) > 1 {
		return
	}
		fmt.Println("No module selected")
		fmt.Println("Usage: nerd-cli -help for more information")
	os.Exit(1)
}

func main() {
	modules := make(map[string]*bool)
	modules["config"] = flag.Bool("config", false, "Module to start the config")
	modules["service"] = flag.Bool("service", false, "Module to start the service")
	help := flag.Bool("help", false, "Show help message")
	initEmpty()
	flagsMap := helpers.GetOptionFlags()
	flag.Parse()
	if *help {
		fmt.Println(services.HelpMessage())
		os.Exit(0)
	}
	module, err := services.GetModule(modules, flagsMap)
	if err != nil {
		panic(err)
	}
	module.GetOptions()
	module.Run()
}
