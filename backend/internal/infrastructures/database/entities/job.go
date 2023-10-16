package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
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

func (e *Job) ToModel() *models.Job {
	return &models.Job{
		ID:         e.ID,
		Name:       e.Name,
		Status:     e.Status,
		CreatedAt:  e.CreatedAt.Unix(),
		UpdatedAt:  e.UpdatedAt.Unix(),
		InputFiles: nil,
	}
}

type Jobs []*Job

func (e *Jobs) ToModel() models.Jobs {
	jobs := make([]*models.Job, 0)

	for _, j := range *e {
		jobs = append(jobs, j.ToModel())
	}

	return jobs
}
