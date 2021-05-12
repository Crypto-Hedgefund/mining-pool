/*
	config.go

	uses `spf13/viper` package to read from config.yml
	populates config and nested structs to allow program to access vars
*/

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Redis    RedisConfig
	Database DatabaseConfig
}

type RedisConfig struct {
	Address string
	Port    int
}

type DatabaseConfig struct {
	User     string
	Password string
}

func Setup() (Config, error) {
	// Config file name
	viper.SetConfigName("config")

	// Path to config
	viper.AddConfigPath("./")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return config, nil
}
