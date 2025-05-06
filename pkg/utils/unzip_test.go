package utils_test

import (
	"archive/zip"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DinnieJ/njav-downloader/pkg/utils"
)

func createTestZip(zipPath string, files map[string]string) error {
	zipfile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	writer := zip.NewWriter(zipfile)
	defer writer.Close()

	for name, content := range files {
		f, err := writer.Create(name)
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(content))
		if err != nil {
			return err
		}
	}
	return nil
}

func TestUnzip(t *testing.T) {
	tmpDir := t.TempDir()

	zipPath := filepath.Join(tmpDir, "test.zip")
	destDir := filepath.Join(tmpDir, "unzipped")

	files := map[string]string{
		"file1.txt":           "Hello, World!",
		"dir1/file2.txt":      "Hello from dir1!",
		"dir1/dir2/file3.txt": "Nested Hello!",
	}

	// Create a test ZIP file
	err := createTestZip(zipPath, files)
	if err != nil {
		t.Fatalf("Failed to create test zip: %v", err)
	}

	// Call the unzip function
	err = utils.Unzip(zipPath, destDir)
	if err != nil {
		t.Fatalf("Unzip failed: %v", err)
	}

	// Check each file
	for name, content := range files {
		path := filepath.Join(destDir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read extracted file %s: %v", name, err)
			continue
		}
		if strings.TrimSpace(string(data)) != content {
			t.Errorf("Content mismatch for %s: expected %q, got %q", name, content, data)
		}
	}
}
