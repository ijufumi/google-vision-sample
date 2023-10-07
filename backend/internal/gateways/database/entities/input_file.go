package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/gateways/database/entities/enums"
	"gorm.io/gorm"
)

type InputFile struct {
	gorm.Model
	ID          string
	JobID       string
	PageNo      uint
	FileKey     string
	FileName    string
	ContentType enums.ContentType
	Size        uint
	Width       uint
	Height      uint
	Status      enums.InputFileStatus
	OutputFiles []*OutputFile
}
