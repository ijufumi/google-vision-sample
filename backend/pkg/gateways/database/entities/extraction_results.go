package entities

import (
	"gorm.io/gorm"
)

type ExtractionResults struct {
	gorm.Model
	ID       string
	Status   string
	ImageUri string
}
