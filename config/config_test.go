package config

import (
	"os"
	"testing"
)

type loggerFake struct{}

func (l *loggerFake) Info(string) {}

func TestGetConfig(t *testing.T) {
	fake := &loggerFake{}
	config, err := GetConfig(fake, ".env")
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != "8080" {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
}

func TestGetConfigNoEnvFile(t *testing.T) {
	fake := &loggerFake{}
	os.Setenv("SERVER_ENABLED", "true")
	os.Setenv("SERVER_PORT", "8080")

	config, err := GetConfig(fake, ".notExists")
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
	os.Setenv("SERVER_ENABLED", "true")

	_, err := GetConfig(fake, ".notExists")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	os.Clearenv()
}
