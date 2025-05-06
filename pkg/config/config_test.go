package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DinnieJ/njav-downloader/pkg/config"
)

func TestGetConfig(t *testing.T) {
	tmpDir := os.TempDir()
	config := &config.Config{
		FolderPath: tmpDir,
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "settings.json"), []byte("{\"testdata\": 123123}"), 0755); err != nil {
		t.Fatalf("Failed to create test config file")
	}

	if err := config.Init(); err != nil {
		t.Fatalf("Failed to init config")
	}
	if config.Get("testData").(float64) != float64(123123) {
		t.Errorf("wrong test data")
	}
	os.Remove(filepath.Join(tmpDir, "settings.json"))
}
