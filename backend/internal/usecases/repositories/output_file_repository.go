package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/models/entities"
	"gorm.io/gorm"
)

type OutputFileRepository interface {
	GetByJobID(db *gorm.DB, jobID string) ([]*entities.OutputFile, error)
	Create(db *gorm.DB, entity ...*entities.OutputFile) error
	DeleteByJobID(db *gorm.DB, jobID string) error
	DeleteByInputFileID(db *gorm.DB, inputFileID string) error
}
