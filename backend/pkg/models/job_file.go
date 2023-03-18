package models

type JobFile struct {
	ID          string `json:"id"`
	JobID       string `json:"jobID"`
	IsOutput    bool   `json:"isOutput"`
	FileKey     string `json:"fileKey"`
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
	Size        int64  `json:"size"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}
