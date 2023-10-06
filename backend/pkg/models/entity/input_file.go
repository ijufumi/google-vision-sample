package entity

import "github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"

type InputFile struct {
	ID          string            `json:"id"`
	JobID       string            `json:"jobID"`
	FileKey     string            `json:"fileKey"`
	FileName    string            `json:"fileName"`
	ContentType enums.ContentType `json:"contentType"`
	Size        uint              `json:"size"`
	CreatedAt   int64             `json:"createdAt"`
	UpdatedAt   int64             `json:"updatedAt"`
	OutputFiles []*OutputFile     `json:"outputFiles"`
}
