package entities

import "gorm.io/gorm"

type OutputFile struct {
	gorm.Model
	ID             string
	JobID          string
	InputFileID    string
	FileKey        string
	FileName       string
	ContentType    string
	Size           uint
	Width          uint
	Height         uint
	ExtractedTexts []*ExtractedText
}
