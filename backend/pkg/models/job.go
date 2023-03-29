package models

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
)

type Job struct {
	ID        string          `json:"id"`
	Status    enums.JobStatus `json:"status"`
	CreatedAt int64           `json:"createdAt"`
	UpdatedAt int64           `json:"updatedAt"`
	JobFiles  []JobFile       `json:"jobFiles"`
}
