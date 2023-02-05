package models

import "time"

type ExtractionResult struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	ImageUri  string    `json:"image_uri"`
	OutputUri *string   `json:"output_uri"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
