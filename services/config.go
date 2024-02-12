package services

import (
	"fmt"
	"nerd-cli/helpers"
)

type Command interface {
	Run()
}

type CommandService struct {
	command helpers.Option
}

func (c *CommandService) Run() {
	fmt.Println("Running Command Service")
}

type ConfigService struct {
	Options  []helpers.Option
	Flags    map[string]*helpers.Option
	AllFlags map[string]interface{}
	module   string
}

func (c *ConfigService) GetOptions() {
	jsonOptions := helpers.OptionParser("flags.json")
	c.Flags = FilterByModule(jsonOptions, c.module)
}



func (c *ConfigService) Run() {
	fmt.Println("Running Config Service")
	FillValues(c.AllFlags, c.Flags)
	fatherCommands := helpers.GetCustomFlags(c.Flags, func(option helpers.Option) bool {
		return GetGroupFathers(option, c.module)
	})
	fatherCommand,error := FilterTrueOptions(fatherCommands)
	if error != nil {
		panic(error)
	}
	fmt.Println(fatherCommand[0])
}


func NewConfigService(flags map[string]interface{}) *ConfigService {
	return &ConfigService{
		AllFlags: flags,
		module:   "config",
	}
}
