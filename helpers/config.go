package helpers

import "github.com/spf13/viper"

type Config struct {
	API_URL  string `mapstructure:"apiUrl"`
	TOKEN  string `mapstructure:"token"`
}

// loadConfig loads the configuration from a file
func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// saveConfig saves the configuration to a file
func SaveConfig(config *Config) error {
	for key, value := range map[string]interface{}{
		"apiUrl": config.API_URL,
		"token": config.TOKEN,
	} {
		viper.Set(key, value)
	}

	return viper.WriteConfig()
}

func SaveConfigValue(key string, value string) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}