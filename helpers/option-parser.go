package helpers

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
)

type FilterCallback func(option Option) bool

type Option struct {
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Default     any    `json:"default"`
	Module      string `json:"module"`
	Value       any    `json:"value"`
	IsFather    bool   `json:"isFather"`
	Fathers 	 []string `json:"fathers"`
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
			value,_ := strconv.Atoi(option.Default.(string))
			flagValue = flag.Int(option.Flag, value, option.Description)
		case "string":
			flagValue = flag.String(option.Flag, option.Default.(string), option.Description)
		case "boolean":
			flagValue = flag.Bool(option.Flag, option.Default.(bool), option.Description)
		}
		flagMap[option.Flag] = flagValue
	}
	return flagMap
}


func GetTrueFlags(flags map[string]*Option) map[string]*Option {
	trueFlags := make(map[string]*Option)

	for key, option := range flags {
		if boolValue, ok := option.Value.(*bool); ok && boolValue != nil && *boolValue {
			trueFlags[key] = option
		}
	}

	return trueFlags
}

func GetFlags (flags map[string]*Option,keys []string) map[string]*Option {
	flagsFound := make(map[string]*Option)
	for _, key := range keys {
		flagsFound[key] = flags[key]
	}
	return flagsFound
}

func GetCustomFlags (flags map[string]*Option,callback FilterCallback) map[string]Option {
	flagsFound := make(map[string]Option)
	for key, option := range flags {
		if callback(*option) {
			flagsFound[key] = *option
		}
	}
	return flagsFound
}

func GetValue(option *Option) interface{} {
	if option.Value == nil {
		return nil
	}
	switch option.Type {
	case "int":
		if value, ok := option.Value.(*int); ok {
			return *value
		}
	case "string":
		if value, ok := option.Value.(*string); ok {
			return *value
		}
	case "boolean":
		if value, ok := option.Value.(*bool); ok {
			return *value
		}
	}

	return nil
}