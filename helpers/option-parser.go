package helpers

import (
	"encoding/json"
	"flag"
	"os"
)

type Option struct {
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Default     any    `json:"default"`
	Module      string `json:"module"`
	Value       any    `json:"value"`
}

func OptionParser(path string) []Option {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var options []Option
	err = json.Unmarshal(file, &options)
	if err != nil {
		panic(err)
	}
	return options
}

func GetOptionFlags() map[string]interface{} {
	flags := OptionParser("flags.json")

	flagMap := make(map[string]interface{})

	for _, option := range flags {
		var flagValue interface{}
		switch option.Type {
		case "int":
			flagValue = flag.Int(option.Flag, option.Default.(int), option.Description)
		case "string":
			flagValue = flag.String(option.Flag, option.Default.(string), option.Description)
		case "boolean":
			flagValue = flag.Bool(option.Flag, option.Default.(bool), option.Description)
		}
		flagMap[option.Name] = flagValue
	}
	return flagMap
}
