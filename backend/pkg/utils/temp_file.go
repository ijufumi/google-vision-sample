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
	subDir := NewRandomString(10)
	dir := fmt.Sprintf("%s/%s", tempDir, subDir)
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(fmt.Sprintf("%s/%s", dir, fileName))
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
