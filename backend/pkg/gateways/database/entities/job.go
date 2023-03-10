package entities

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID             string
	Status         enums.JobStatus
	ExtractedTexts []ExtractedText
	Files          []JobFile
}
