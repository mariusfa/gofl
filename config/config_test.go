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
	if config.AppName != "testapp" {
		t.Errorf("expected app name to be test, got %v", config.AppName)
	}
}

func TestGetConfigNoEnvFile(t *testing.T) {
	fake := &loggerFake{}
	os.Setenv("PORT", "8080")
	os.Setenv("APP_NAME", "testapp")

	config, err := GetConfig(fake, ".notExists")
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != "8080" {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
	if config.AppName != "testapp" {
		t.Errorf("expected app name to be test, got %v", config.AppName)
	}
	os.Clearenv()
}

func TestGetConfigMissingPort(t *testing.T) {
	fake := &loggerFake{}
	os.Setenv("APP_NAME", "testapp")

	_, err := GetConfig(fake, ".notExists")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	os.Clearenv()
}
