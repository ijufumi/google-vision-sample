package entities

import "gorm.io/gorm"

type File struct {
	gorm.Model
	ID                 string
	ExtractionResultID string
	IsOutput           bool
	FileKey            string
	FileName           string
	ContentType        string
	Size               int32
}
