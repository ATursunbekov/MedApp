package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Port int
	}

	Mongo struct {
		URI      string
		Database string
	}

	Redis struct {
		Address  string
		Password string
	}
}

var AppConfig Config

func LoadYAMLConfig(path string) {
	viper.SetConfigFile(path)   // e.g., config.yml
	viper.SetConfigType("yaml") // or "yml"

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
