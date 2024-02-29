package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

type LoggerInterface interface {
	Info(string)
}

func GetConfig(logger LoggerInterface) (*Config, error) {
	defaultViper := viper.New()

	defaultViper.SetConfigName("application")
	defaultViper.SetConfigType("yaml")
	defaultViper.AddConfigPath(".")

	defaultViper.AutomaticEnv()
	defaultViper.SetEnvPrefix("")

	if err := defaultViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info("Config file not found; using defaults")
		} else {
			return nil, err
		}
	}

	var config Config
	err := defaultViper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	err = validateConfig(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	if config.Port == 0 {
		return errors.New("port is required")
	}

	return nil
}
