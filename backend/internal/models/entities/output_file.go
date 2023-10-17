package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
)

type OutputFile struct {
	ID             string            `json:"id"`
	JobID          string            `json:"jobID"`
	InputFileID    string            `json:"inputFileID"`
	FileKey        string            `json:"fileKey"`
	FileName       string            `json:"fileName"`
	ContentType    enums.ContentType `json:"contentType"`
	Size           uint              `json:"size"`
	CreatedAt      int64             `json:"createdAt"`
	UpdatedAt      int64             `json:"updatedAt"`
	ExtractedTexts []*ExtractedText  `json:"extractedTexts"`
}

type OutputFiles []*OutputFile
