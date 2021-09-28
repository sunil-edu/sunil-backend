package utils

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DbDriver     string `mapstructure:"DB_DRIVER"`
	DbSource     string `mapstructure:"DB_SOURCE"`
	ServerAddr   string `mapstructure:"SERVER_ADDRESS"`
	AllowOrigins string `mapstructure:"ALLOW_ORIGINS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
