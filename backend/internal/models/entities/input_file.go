package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
)

type InputFile struct {
	ID          string                `json:"id"`
	JobID       string                `json:"jobID"`
	PageNo      uint                  `json:"pageNo"`
	FileKey     string                `json:"fileKey"`
	FileName    string                `json:"fileName"`
	ContentType enums.ContentType     `json:"contentType"`
	Size        uint                  `json:"size"`
	Width       uint                  `json:"width"`
	Height      uint                  `json:"height"`
	Status      enums.InputFileStatus `json:"status"`
	CreatedAt   int64                 `json:"createdAt"`
	UpdatedAt   int64                 `json:"updatedAt"`
	OutputFiles OutputFiles           `json:"outputFiles"`
}

type InputFiles []*InputFile
