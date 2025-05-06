package driver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func getDownloadDriverPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	absPath := filepath.Join(homeDir, DriverDirectory, "DinnieJ")
	return absPath, nil
}

func createFolderIfNotExist(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if os.IsNotExist(err) || !info.IsDir() {
		os.MkdirAll(path, 0755)
	}
	return nil
}

func checkFileExist(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if os.IsNotExist(err) || info.IsDir() {
		return fmt.Errorf("file not exist in provided path")
	}
	return nil
}
func downloadDriver(dest string) error {
	resp, err := http.Get(DriverDownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to download driver")
	}

	binary, err := io.ReadAll(resp.Body)
	if err := os.WriteFile(dest, binary, 0775); err != nil {
		return err
	}

	return nil
}

func GetDriverPath() (string, error) {
	absDriverFolderPath, err := getDownloadDriverPath()
	createFolderIfNotExist(absDriverFolderPath) // creating folder regardless
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(absDriverFolderPath, OutputFilename)
	zipPath := filepath.Join(absDriverFolderPath, "driver.zip")
	if err := checkFileExist(fullPath); err != nil {
		if errDownload := downloadDriver(zipPath); errDownload != nil {
			return "", errDownload
		}
		unzip(zipPath, absDriverFolderPath)
	}

	return fullPath, nil
}
