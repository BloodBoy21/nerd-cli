package commands

import (
	"fmt"
	"nerd-cli/helpers"
)
type CLICommand interface {
	Run()
}

type CommandService struct {
	Command helpers.Option
	Flags  map[string] helpers.Option
}

func (c *CommandService) Run() {
	fmt.Printf("Running %s command\n", c.Command.Flag)
	for key, value := range c.Flags {
		fmt.Printf("Flag: %s, Value: %v\n", key, helpers.GetValue(&value))
	}
}