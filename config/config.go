package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	AppName string
}

type LoggerInterface interface {
	Info(string)
}

func GetConfig(logger LoggerInterface, filename string) (*Config, error) {
	err := godotenv.Load(filename)
	if err != nil {
		logger.Info("Could not load .env file. Using ENV variables")
	}

	config := &Config{
		Port:    os.Getenv("PORT"),
		AppName: os.Getenv("APP_NAME"),
	}

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	if config.Port == "" {
		return errors.New("env var PORT is required")
	}
	if config.AppName == "" {
		return errors.New("env var APP_NAME is required")
	}

	return nil
}
