package utils

import (
	"fmt"
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
