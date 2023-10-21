package repositories

import (
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	"gorm.io/gorm"
)

type ExtractedTextRepository interface {
	GetByID(db *gorm.DB, id string) (*models.ExtractedText, error)
	Create(db *gorm.DB, entity ...*models.ExtractedText) error
	DeleteByOutputFileID(db *gorm.DB, outputFileID string) error
}
