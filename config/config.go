package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
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
		Port:    os.Getenv("SERVER_PORT"),
	}

	err = validateServerConfig(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func validateServerConfig(config *Config) error {
	if os.Getenv("SERVER_ENABLED") != "true" {
		return nil
	}
	if config.Port == "" {
		return errors.New("env var SERVER_PORT is required")
	}

	return nil
}
