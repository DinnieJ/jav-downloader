package utils

import "os"

func CheckFileExist(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func CheckFolderExist(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func CreateFolderIfNotExist(path string) error {
	if CheckFolderExist(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}
