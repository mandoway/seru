package files

import (
	"io"
	"os"
	"path"
)

// Copy copies src to dst like the cp command.
func Copy(src, dst string) error {
	if dst == src {
		return os.ErrInvalid
	}

	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcF *os.File) {
		_ = srcF.Close()
	}(srcF)

	info, err := srcF.Stat()
	if err != nil {
		return err
	}

	dstF, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer func(dstF *os.File) {
		_ = dstF.Close()
	}(dstF)

	if _, err := io.Copy(dstF, srcF); err != nil {
		return err
	}
	return nil
}

func AddFolderToFilePath(inputFilePath, folder string) string {
	dir, file := path.Split(inputFilePath)
	return path.Join(dir, folder, file)
}
