package storage

import (
	"os"
	"testing"
)

func TestSaveAndLoadConfig(t *testing.T) {
	// Test saving and loading
	// 1. Create a test config
	cfg := Config{
		Salt:          "12345efrerff23445",
		SetupComplete: true,
	}

	// 2. Save it to "test_config.json"
	err := SaveConfig("test_config.json", &cfg)
	if err != nil {
		t.Error(err)
	}
	// 3. Load it back
	test_cfg, err := LoadConfig("test_config.json")
	if err != nil {
		t.Error(err)
	}
	// 4. Check Salt and SetupComplete match
	if test_cfg.Salt != cfg.Salt {
		t.Errorf("Salt mismatch: wanted %s, got %s", cfg.Salt, test_cfg.Salt)
	}
	if test_cfg.SetupComplete != cfg.SetupComplete {
		t.Errorf("SetupComplete mismatch: wanted %v, got %v", cfg.SetupComplete, test_cfg.SetupComplete)
	}
	// 5. Clean up: os.Remove("test_config.json")
	os.Remove("test_config.json")
}

func TestLoadConfigNotExists(t *testing.T) {
	// Test loading when file doesn't exist
	// Check if file exists
	// 1. Call LoadConfig on nonexistent file
	cfg, err := LoadConfig("nonexistent.json")

	// 2. Check that cfg is nil and err is nil (first run case)
	if cfg != nil {
		t.Errorf("Expected nil config, got %v", cfg)
	}
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}
