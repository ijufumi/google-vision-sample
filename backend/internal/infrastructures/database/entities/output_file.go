package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	"gorm.io/gorm"
)

type OutputFile struct {
	gorm.Model
	ID             string
	JobID          string
	InputFileID    string
	FileKey        string
	FileName       string
	ContentType    enums.ContentType
	Size           uint
	Width          uint
	Height         uint
	ExtractedTexts []*ExtractedText
}
