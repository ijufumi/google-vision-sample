package utils

import (
	"os"
)

func NewTempFile() (*os.File, error) {
	return NewTempFileWithName(NewRandomString(10))
}

func NewTempFileWithName(fileName string) (*os.File, error) {
	tempDir := os.TempDir()
	file, err := os.CreateTemp(tempDir, fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
