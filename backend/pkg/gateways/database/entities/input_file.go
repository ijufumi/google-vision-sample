package entities

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"gorm.io/gorm"
)

type InputFile struct {
	gorm.Model
	ID          string
	JobID       string
	PageNo      uint
	FileKey     string
	FileName    string
	ContentType string
	Size        uint
	Width       uint
	Height      uint
	Status      enums.InputFileStatus
	OutputFiles []*OutputFile
}
