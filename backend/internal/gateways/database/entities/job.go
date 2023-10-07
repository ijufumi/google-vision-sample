package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/gateways/database/entities/enums"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID              string
	Name            string
	OriginalFileKey string
	Status          enums.JobStatus
	InputFiles      []InputFile
}
