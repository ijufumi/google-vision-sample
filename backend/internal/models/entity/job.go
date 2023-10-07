package entity

import (
	"github.com/ijufumi/google-vision-sample/internal/gateways/database/entities/enums"
)

type Job struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Status     enums.JobStatus `json:"status"`
	CreatedAt  int64           `json:"createdAt"`
	UpdatedAt  int64           `json:"updatedAt"`
	InputFiles []InputFile     `json:"inputFiles"`
}
