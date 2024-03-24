package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(filename string) *viper.Viper {
	config := viper.New()

	config.SetConfigName(filename)
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s\n", err.Error())
	}

	return config
}
