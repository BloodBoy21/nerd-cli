package services

import (
	"nerd-cli/commands"
	"nerd-cli/helpers"
)



type CLIService struct {
	Options  []helpers.Option
	Flags    map[string]*helpers.Option
	AllFlags map[string]interface{}
	module   string
}

func (c *CLIService) GetOptions() {
	jsonOptions := helpers.OptionParser("flags.json")
	c.Flags = FilterByModule(jsonOptions, c.module)
}



func (c *CLIService) Run() {
	FillValues(c.AllFlags, c.Flags)
	fatherCommands := helpers.GetCustomFlags(c.Flags, func(option helpers.Option) bool {
		return GetGroupFathers(option, c.module)
	})
	fatherCommand,error := FilterTrueOptions(fatherCommands)
	if error != nil {
		panic(error)
	}
	commandFlags := helpers.GetCustomFlags(c.Flags, func(option helpers.Option) bool {
		return GetGroupChildren(option, fatherCommand.Flag)
	})
	command := &commands.CommandService{
		Command: fatherCommand,
		Flags:   commandFlags,
	}

	command.Run()
}


func NewService(flags map[string]interface{},module string) *CLIService {
	return &CLIService{
		AllFlags: flags,
		module:   module,
	}
}
