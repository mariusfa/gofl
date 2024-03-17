package config

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

// Example of config to be sent into GetConfig
// type Config struct {
// 	ServerEnabled string `required`
// 	ServerPort string `required ServerEnabled`
// }

type LoggerInterface interface {
	Info(string)
}

func GetConfig(logger LoggerInterface, filename string, config any) error {
	err := godotenv.Load(filename)
	if err != nil {
		logger.Info("Could not load .env file. Using ENV variables")
	}

	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		fieldName := fieldInfo.Name
		envName := getEnvName(fieldName)
		envValue := os.Getenv(envName)
		requiredTag := fieldInfo.Tag.Get("required")

		if envValue == "" && requiredTag != "false" {
			return errors.New("env var " + envName + " is required")
		}

		v.Field(i).SetString(envValue)
	}
	return nil
}

// Gets env name from field name
// Example: ServerPort -> SERVER_PORT
func getEnvName(fieldName string) string {
	envName := ""
	for i, letter := range fieldName {
		if i > 0 && letter >= 'A' && letter <= 'Z' {
			envName += "_"
		}
		envName += string(letter)
	}
	envName = strings.ToUpper(envName)
	return envName
}
