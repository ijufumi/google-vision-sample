package entities

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"gorm.io/gorm"
)

type ExtractionResult struct {
	gorm.Model
	ID             string
	Status         enums.ExtractionResultStatus
	ImageKey       string
	OutputKey      *string
	ExtractedTexts []ExtractedText
}
