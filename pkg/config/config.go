package config

import (
	"os"
	"path/filepath"

	"github.com/DinnieJ/njav-downloader/pkg/utils"
	"github.com/spf13/viper"
)

type Config struct {
	FolderPath string
}

func (c *Config) Init() error {
	if c.FolderPath == "" {
		homeDir, _ := os.UserHomeDir()
		c.FolderPath = filepath.Join(homeDir, CONFIG_FOLDER)
	}
	cnfFilePath := filepath.Join(c.FolderPath, CONFIG_FILENAME)
	if !utils.CheckFolderExist(c.FolderPath) {
		if err := os.MkdirAll(c.FolderPath, os.ModePerm); err != nil {
			return err
		}
	}
	if !utils.CheckFileExist(cnfFilePath) {
		if err := c.createDefaultConfig(); err != nil {
			return err
		}
	}
	viper.AddConfigPath(c.FolderPath)
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType("json")
	return viper.ReadInConfig()
}

func (c *Config) Get(key string) any {
	return viper.Get(key)
}

func (c *Config) createDefaultConfig() error {
	defaultConfigStr := "{\"maximumThread\": 8}\n"
	if err := os.WriteFile(filepath.Join(c.FolderPath, CONFIG_FILENAME), []byte(defaultConfigStr), 0755); err != nil {
		return err
	}
	return nil
	// return nil
}
