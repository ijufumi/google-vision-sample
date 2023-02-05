package utils

import (
	"os"
)

func NewTempFile() (*os.File, error) {
	tempDir := os.TempDir()
	fileName := NewRandomString(10)
	file, err := os.CreateTemp(tempDir, fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
