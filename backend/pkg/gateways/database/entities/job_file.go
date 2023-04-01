package entities

import "gorm.io/gorm"

type JobFile struct {
	gorm.Model
	ID             string
	JobID          string
	IsOutput       bool
	FileKey        string
	FileName       string
	ContentType    string
	Size           uint
	Width          uint
	Height         uint
	ExtractedTexts []*ExtractedText
}
