package services

import (
	"errors"
	"nerd-cli/helpers"
)

type Service interface {
	GetOptions()
	Run()
}

func filterTrueModules(modules map[string]*bool) []string {
	var trueModules []string
	for key, value := range modules {
		if *value {
			trueModules = append(trueModules, key)
		}
	}
	return trueModules
}

func FilterByModule(modules []helpers.Option,module string) map[string]*helpers.Option {
	filteredOptions := make(map[string]*helpers.Option)
	for _, option := range modules {
		if option.Module == module {
			filteredOptions[option.Name] = &option
		}
	}
	return filteredOptions
}



func GetModule(modules map[string]*bool,flags map[string]interface{}) (Service, error) {
	if modules["config"] == nil && modules["service"] == nil {
		return nil, errors.New("No module selected")
	}
	if len(filterTrueModules(modules)) > 1 {
		return nil, errors.New("Multiple modules selected")
	}
	selectedModule := filterTrueModules(modules)[0]
	
	switch selectedModule {
	case "config":
		return NewConfigService(flags), nil
	}
	return nil, errors.New("No module selected")
}

func FillValues(flags map[string]interface{}, options map[string]*helpers.Option) {
	for key, option := range options {
		switch option.Type {
		case "int":
			option.Value = flags[key].(*int)
		case "string":
			option.Value = flags[key].(*string)
		case "boolean":
			option.Value = flags[key].(*bool)
		}
	}
}

func GetValue(option *helpers.Option) interface{} {
	switch option.Type {
	case "int":
		return *option.Value.(*int)
	case "string":
		return *option.Value.(*string)
	case "boolean":
		return *option.Value.(*bool)
	}
	return nil
}