package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestLoad(t *testing.T) {
	viper.Reset()
	
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.FrameworksDir == "" {
		t.Error("FrameworksDir should not be empty")
	}
	if cfg.SessionFile == "" {
		t.Error("SessionFile should not be empty")
	}
	if cfg.OutputFormat == "" {
		t.Error("OutputFormat should not be empty")
	}
}

func TestLoadWithEnvVars(t *testing.T) {
	viper.Reset()
	
	os.Setenv("TINYTHOUGHTS_FRAMEWORKS_DIR", "/test/frameworks")
	os.Setenv("TINYTHOUGHTS_SESSION_FILE", "/test/session.json")
	os.Setenv("TINYTHOUGHTS_OUTPUT_FORMAT", "text")
	defer func() {
		os.Unsetenv("TINYTHOUGHTS_FRAMEWORKS_DIR")
		os.Unsetenv("TINYTHOUGHTS_SESSION_FILE")
		os.Unsetenv("TINYTHOUGHTS_OUTPUT_FORMAT")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.FrameworksDir != "/test/frameworks" {
		t.Errorf("Expected '/test/frameworks', got '%s'", cfg.FrameworksDir)
	}
	if cfg.SessionFile != "/test/session.json" {
		t.Errorf("Expected '/test/session.json', got '%s'", cfg.SessionFile)
	}
	if cfg.OutputFormat != "text" {
		t.Errorf("Expected 'text', got '%s'", cfg.OutputFormat)
	}
}

func TestLoadDefaults(t *testing.T) {
	viper.Reset()
	
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.OutputFormat != "json" {
		t.Errorf("Expected default output format 'json', got '%s'", cfg.OutputFormat)
	}
}
