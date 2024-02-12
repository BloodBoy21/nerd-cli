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
	FillValues(c.AllFlags, c.Flags)
	for key, option := range c.Flags {
		fmt.Println(key, option)
		value := GetValue(option)
		fmt.Println(value)
	}
}

func NewConfigService(flags map[string]interface{}) *ConfigService {
	return &ConfigService{
		AllFlags: flags,
		module:   "config",
	}
}
