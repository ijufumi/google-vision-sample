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
	InputFiles      InputFiles
}

func (e *Job) ToModel() *models.Job {
	return &models.Job{
		ID:              e.ID,
		Name:            e.Name,
		Status:          e.Status,
		OriginalFileKey: e.OriginalFileKey,
		CreatedAt:       e.CreatedAt.Unix(),
		UpdatedAt:       e.UpdatedAt.Unix(),
		InputFiles:      e.InputFiles.ToModel(),
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

func FromJobModel(jobModel *models.Job) *Job {
	return &Job{
		ID:              jobModel.ID,
		Name:            jobModel.Name,
		Status:          jobModel.Status,
		OriginalFileKey: jobModel.OriginalFileKey,
		InputFiles:      FromInputFilesModel(jobModel.InputFiles),
	}
}
