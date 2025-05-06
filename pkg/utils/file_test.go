package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckFileExist(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")

	// File does not exist yet
	if CheckFileExist(tmpFile) {
		t.Errorf("Expected CheckFileExist to return false for non-existent file")
	}

	// Create the file
	err := os.WriteFile(tmpFile, []byte("hello"), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Now file should exist
	if !CheckFileExist(tmpFile) {
		t.Errorf("Expected CheckFileExist to return true for existing file")
	}
}

func TestCheckFolderExist(t *testing.T) {
	tmpDir := t.TempDir()
	subFolder := filepath.Join(tmpDir, "folder")

	// Folder does not exist yet
	if CheckFolderExist(subFolder) {
		t.Errorf("Expected CheckFolderExist to return false for non-existent folder")
	}

	// Create the folder
	err := os.Mkdir(subFolder, 0755)
	if err != nil {
		t.Fatalf("Failed to create folder: %v", err)
	}

	// Now folder should exist
	if !CheckFolderExist(subFolder) {
		t.Errorf("Expected CheckFolderExist to return true for existing folder")
	}
}

func TestCreateFolderIfNotExist(t *testing.T) {
	tmpDir := t.TempDir()
	newFolder := filepath.Join(tmpDir, "newfolder")

	// Folder should not exist
	if CheckFolderExist(newFolder) {
		t.Fatalf("Folder should not exist before creation")
	}

	// Create the folder
	err := CreateFolderIfNotExist(newFolder)
	if err != nil {
		t.Fatalf("CreateFolderIfNotExist failed: %v", err)
	}

	// Folder should now exist
	if !CheckFolderExist(newFolder) {
		t.Errorf("Expected folder to exist after CreateFolderIfNotExist")
	}

	// Calling it again should not return an error
	if err := CreateFolderIfNotExist(newFolder); err != nil {
		t.Errorf("CreateFolderIfNotExist should not fail on existing folder")
	}
}
