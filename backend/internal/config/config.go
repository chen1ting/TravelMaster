package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DBHost string `mapstructure:"db_host"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
	DBName string `mapstructure:"db_name"`
}

var Conf Config

func InitConfig() {// hardcoded to read from same dir :(, probably can read in from env if we have time to refactor
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
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}