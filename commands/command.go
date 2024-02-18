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


func getCommandCallback(command helpers.Option) func(*CommandService) {
	switch command.Flag {
	case "login":
		return Login
	default:
		return nil
	}
}

func (command *CommandService) Run() {
	callback := getCommandCallback(command.Command)
	if callback == nil {
		fmt.Println("Command not found")
		return
	}
	callback(command)
}