package entities

import "gorm.io/gorm"

type ExtractedText struct {
	gorm.Model
	ID     string
	JobID  string
	Text   string
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}
