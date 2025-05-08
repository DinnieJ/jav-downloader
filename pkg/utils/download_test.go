package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DinnieJ/njav-downloader/pkg/utils"
)

func TestDownload(t *testing.T) {
	tmpDir := os.TempDir()
	testFileUrl := "https://link.testfile.org/15MB"
	dest := filepath.Join(tmpDir, "testData.zip")

	if err := utils.DownloadFile(testFileUrl, dest); err != nil {
		t.Errorf("failed to download file")
	}
	os.Remove(dest)
}
