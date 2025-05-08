package driver

import (
	"os"
	"path/filepath"

	"github.com/DinnieJ/njav-downloader/pkg/config"
	"github.com/DinnieJ/njav-downloader/pkg/utils"
)

func GetDriverPath(cfg *config.Config) (string, error) {
	driverPath := filepath.Join(cfg.FolderPath, OutputFilename)
	if !utils.CheckFileExist(driverPath) {
		LOGGER.Info("Web Driver not found, downloading....")
		if err := DownloadDriver(cfg.FolderPath); err != nil {
			return "", err
		}
	}
	return driverPath, nil
}

func DownloadDriver(dest string) error {
	tmpDir := os.TempDir()
	zipPath := filepath.Join(tmpDir, "driver.zip")
	if err := utils.DownloadFile(DriverDownloadUrl, zipPath); err != nil {
		return err
	}
	if err := utils.Unzip(zipPath, dest); err != nil {
		return err
	}
	os.Remove(zipPath)
	return nil
}
