package services

import (
	"errors"
	"nerd-cli/helpers"
)

type Service interface {
	GetOptions()
	Run()
}
func FilterTrueOptions(options map[string]*helpers.Option) ([]*helpers.Option,error) {
	var trueOption []*helpers.Option
	for _, option := range options {
		if *option.Value.(*bool) {
			trueOption = append(trueOption, option)
		}
	}
	if len(trueOption) == 0 {
		return nil,errors.New("No option selected")
	}
	return trueOption,nil
}

func FilterTrueModule(options map[string]*bool) []string {
	var trueOption []string
	for key, value := range options {
		if *value {
			trueOption = append(trueOption, key)
		}
	}
	return trueOption
}

func FilterByModule(modules []helpers.Option, module string) map[string]*helpers.Option {
	filteredOptions := make([]helpers.Option, 0)
	filteredDict := make(map[string]*helpers.Option)
	for _, option := range modules {
		if option.Module == module {
			filteredOptions = append(filteredOptions, option)
		}
	}
	for _, option := range filteredOptions {
		filteredDict[option.Flag] = &option
	}
	return filteredDict
}

func GroupByModule(modules []helpers.Option) map[string][]helpers.Option {
	groupedOptions := make(map[string][]helpers.Option)
	for _, option := range modules {
		groupedOptions[option.Module] = append(groupedOptions[option.Module], option)
	}
	return groupedOptions
}

func HelpMessage() string {
	options := helpers.OptionParser("flags.json")
	message := "Usage: nerd-cli [module] [flags]\n\n"
	message += "Modules:\n--config\n--service\n\n"
	groups := GroupByModule(options)
	message += "Flags:\n"
	for key, value := range groups {
		message += " " + key + ":\n"
		for _, option := range value {
			message += "	--" + option.Flag + " " + option.Name + " " + option.Description + "\n"
		}
		message += "\n"
	}
	return message
}

func GetModule(modules map[string]*bool, flags map[string]interface{}) (Service, error) {
	if modules["config"] == nil && modules["service"] == nil {
		return nil, errors.New("No module selected")
	}
	if len(FilterTrueModule(modules)) > 1 {
		return nil, errors.New("Multiple modules selected")
	}

	if len(FilterTrueModule(modules)) == 0 {
		return nil, errors.New("No module selected")
	}

	selectedModule := FilterTrueModule(modules)[0]

	switch selectedModule {
	case "config":
		return NewConfigService(flags), nil
	}
	return nil, errors.New("No module selected")
}

func FillValues(flags map[string]interface{}, options map[string]*helpers.Option) {
	for _, option := range options {
		key := option.Flag
		switch option.Type {
		case "int":
			if value, ok := flags[key].(*int); ok {
				option.Value = value
			}
		case "string":
			if value, ok := flags[key].(*string); ok {
				option.Value = value
			}
		case "boolean":
			if value, ok := flags[key].(*bool); ok {
				option.Value = value
			}
		}
	}
}

func GetValue(option *helpers.Option) interface{} {
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

func GetGroupFathers(option helpers.Option,module string) bool {
	return option.Module == module && option.IsFather
}
