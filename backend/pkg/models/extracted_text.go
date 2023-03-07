package models

type ExtractedText struct {
	ID                 string  `json:"id"`
	ExtractionResultID string  `json:"extractionResultID"`
	Text               string  `json:"text"`
	Top                float64 `json:"top"`
	Bottom             float64 `json:"bottom"`
	Left               float64 `json:"left"`
	Right              float64 `json:"right"`
	CreatedAt          int64   `json:"createdAt"`
	UpdatedAt          int64   `json:"updatedAt"`
}
