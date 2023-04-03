package entities

import "gorm.io/gorm"

type InputFile struct {
	gorm.Model
	ID          string
	JobID       string
	FileKey     string
	FileName    string
	ContentType string
	Size        uint
	Width       uint
	Height      uint
	OutputFiles []*OutputFile
}
