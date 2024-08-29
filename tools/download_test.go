package tools

import (
	"errors"
	"os"
	"testing"
)

func TestDownloadFromGithub(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	url := "https://github.com/microsoft/TypeScript/releases/download/v5.5.4/typescript-5.5.4.tgz"
	file := "ts"
	err := downloadFile(file, url)
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		t.Error("file does not exist")
	}

	_ = os.Remove(file)
}
