package models

type JobFile struct {
	ID                 string `json:"id"`
	ExtractionResultID string `json:"extractionResultID"`
	IsOutput           bool   `json:"isOutput"`
	FileKey            string `json:"fileKey"`
	FileName           string `json:"fileName"`
	ContentType        string `json:"contentType"`
	Size               int32  `json:"size"`
	CreatedAt          int64  `json:"createdAt"`
	UpdatedAt          int64  `json:"updatedAt"`
}
