package entities

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
)

type ExtractionResult struct {
	BaseEntity
	ID             string
	Status         enums.ExtractionResultStatus
	ImageKey       string
	OutputKey      *string
	ExtractedTexts []ExtractedText
}
