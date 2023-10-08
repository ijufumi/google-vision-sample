package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	"gorm.io/gorm"
)

type InputFileRepository interface {
	GetByID(db *gorm.DB, iD string) (*entities.InputFile, error)
	GetByJobID(db *gorm.DB, jobID string) ([]*entities.InputFile, error)
	Create(db *gorm.DB, entity ...*entities.InputFile) error
	Update(db *gorm.DB, entity *entities.InputFile) error
	DeleteByJobID(db *gorm.DB, jobID string) error
}
