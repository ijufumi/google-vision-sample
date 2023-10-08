package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	"gorm.io/gorm"
)

type ExtractedTextRepository interface {
	GetByID(db *gorm.DB, id string) (*entities.ExtractedText, error)
	Create(db *gorm.DB, entity ...*entities.ExtractedText) error
	DeleteByOutputFileID(db *gorm.DB, outputFileID string) error
}
