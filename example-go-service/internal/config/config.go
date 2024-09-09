package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"env"`
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

var config Config

func Get() Config {
	return config
}

func Load() {
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(err.Error())
		}
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Config loaded")
}
