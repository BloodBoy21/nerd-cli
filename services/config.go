package services

import (
	"fmt"
	"nerd-cli/helpers"
)

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
	for key, option := range c.Flags {
		flagPointer := c.AllFlags[key]
		fmt.Println(key, option)
		switch option.Type {
		case "int":
			fmt.Println(*flagPointer.(*int))
		case "string":

			fmt.Println(*flagPointer.(*string))
		case "boolean":
			fmt.Println(*flagPointer.(*bool))
		}
	}
}

func NewConfigService(flags map[string]interface{}) *ConfigService {
	return &ConfigService{
		AllFlags: flags,
		module:   "config",
	}
}
