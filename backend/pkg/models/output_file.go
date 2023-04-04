package models

type OutputFile struct {
	ID             string           `json:"id"`
	JobID          string           `json:"jobID"`
	InputFileID    string           `json:"inputFileID"`
	FileKey        string           `json:"fileKey"`
	FileName       string           `json:"fileName"`
	ContentType    string           `json:"contentType"`
	Size           uint             `json:"size"`
	CreatedAt      int64            `json:"createdAt"`
	UpdatedAt      int64            `json:"updatedAt"`
	ExtractedTexts []*ExtractedText `json:"extractedTexts"`
}
