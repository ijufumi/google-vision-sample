package entities

import (
	enums2 "github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	"gorm.io/gorm"
)

type InputFile struct {
	gorm.Model
	ID          string
	JobID       string
	PageNo      uint
	FileKey     string
	FileName    string
	ContentType enums2.ContentType
	Size        uint
	Width       uint
	Height      uint
	Status      enums2.InputFileStatus
	OutputFiles []*OutputFile
}
