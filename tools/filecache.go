package tools

import (
	"errors"
	"github.com/mandoway/seru/version"
	"log"
	"net/url"
	"os"
	"path"
)

var versionedDownloadUrl, _ = url.JoinPath("https://github.com/mandoway/seru/releases/download", version.Version)

func EnsureFileExists(filepath string) error {
	_, err := os.Stat(filepath)
	if err == nil {
		return nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	filename := path.Base(filepath)
	downloadPath, err := url.JoinPath(versionedDownloadUrl, filename)
	if err != nil {
		return err
	}

	log.Println("One-time download from", downloadPath, "into", filepath)
	err = downloadFile(filepath, downloadPath)
	if err != nil {
		_ = os.Remove(filepath)
		return err
	}

	return nil
}
