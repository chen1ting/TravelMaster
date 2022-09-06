package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost string `mapstructure:"db_host"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
	DBName string `mapstructure:"db_name"`
	DBPort string `mapstructure:"db_port"`

	SessionRedisHost string `mapstructure:"session_redis_host"`
	SessionRedisPort string `mapstructure:"session_redis_port"`
}

func NewConfig() Config { // hardcoded to read from same dir :(, probably can read in from env if we have time to refactor
	configName := "config_prod"
	if os.Getenv("APP_ENV") == "development" {
		configName = "config_dev.yml"
	}
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName(configName)

	fmt.Printf("initializing config: %s \n", configName)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Sprintf("config file %s not found", configName))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Sprintf("unexpected err, reading config file: %v", err))
		}
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}
