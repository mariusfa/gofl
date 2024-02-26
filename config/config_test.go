package config

import "testing"

func TestGetConfig(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != 8080 {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
}
