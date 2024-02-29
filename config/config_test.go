package config

import "testing"

type loggerFake struct{}

func (l *loggerFake) Info(string) {}

func TestGetConfig(t *testing.T) {
	fake := &loggerFake{}
	config, err := GetConfig(fake)
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != 8080 {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
}
