package entities

import "gorm.io/gorm"

type ExtractedText struct {
	gorm.Model
	ExtractionResultID string
	Text               string
	Top                float64
	Bottom             float64
	Left               float64
	Right              float64
}
