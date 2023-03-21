package utils

import (
	"fmt"
	"io"
	"os"
)

func NewTempFile() (*os.File, error) {
	return NewTempFileWithName(NewRandomString(10))
}

func NewTempFileWithName(fileName string) (*os.File, error) {
	tempDir := os.TempDir()
	file, err := os.Create(fmt.Sprintf("%s/%s", tempDir, fileName))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Copy(src *os.File, dst *os.File) error {
	_, err := src.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	return err
}
