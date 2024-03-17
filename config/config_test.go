package config

import (
	"os"
	"testing"
)

type loggerFake struct{}

func (l *loggerFake) Info(msg string) {
	println(msg)
}

func TestGetConfig(t *testing.T) {
	fake := &loggerFake{}
	type Config struct {
		Port string
	}

	var config Config

	err := GetConfig(fake, ".env", &config)
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != "8080" {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
	os.Clearenv()
}

func TestGetConfigNoEnvFile(t *testing.T) {
	fake := &loggerFake{}
	os.Setenv("PORT", "8080")
	type Config struct {
		Port string
	}

	var config Config

	err := GetConfig(fake, ".notExists", &config)
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != "8080" {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
	os.Clearenv()
}

func TestGetConfigMissingPort(t *testing.T) {
	fake := &loggerFake{}
	type Config struct {
		Port string
	}

	var config Config

	err := GetConfig(fake, ".notExists", &config)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	os.Clearenv()
}
