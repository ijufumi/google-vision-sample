package models

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
)

type ExtractionResult struct {
	ID             string                       `json:"id"`
	Status         enums.ExtractionResultStatus `json:"status"`
	ImageKey       string                       `json:"imageKey"`
	OutputKey      string                       `json:"outputKey"`
	CreatedAt      int64                        `json:"createdAt"`
	UpdatedAt      int64                        `json:"updatedAt"`
	ExtractedTexts []ExtractedText              `json:"extractedTexts"`
}
