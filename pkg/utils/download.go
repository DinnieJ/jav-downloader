package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	ErrFailed       = fmt.Errorf("failed to fetch file")
	ErrFailedStatus = fmt.Errorf("failed to download file due to not ok status")
)

func DownloadFile(url string, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return ErrFailed
	}
	if resp.StatusCode != 200 {
		return ErrFailedStatus
	}
	defer resp.Body.Close()
	binData, err := io.ReadAll(resp.Body)

	if err := os.WriteFile(dest, binData, 0755); err != nil {
		return err
	}

	return nil
}
