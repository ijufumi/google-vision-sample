package entities

import (
	"gorm.io/gorm"
)

type ExtractionResult struct {
	gorm.Model
	ID             string
	Status         string
	ImageUri       string
	ExtractedTexts []ExtractedText
}
