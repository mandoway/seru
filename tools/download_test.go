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
	url := "https://github.com/mandoway/seru/releases/download/v0.0.1-alpha/cue.jar"
	file := "cue.jar"
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
