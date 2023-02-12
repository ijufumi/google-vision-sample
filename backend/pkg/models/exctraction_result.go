package models

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
)

type ExtractionResult struct {
	ID             string                       `json:"id"`
	Status         enums.ExtractionResultStatus `json:"status"`
	ImageUri       string                       `json:"imageUri"`
	OutputUri      string                       `json:"outputUri"`
	CreatedAt      int64                        `json:"createdAt"`
	UpdatedAt      int64                        `json:"updatedAt"`
	ExtractedTexts []ExtractedText              `json:"extractedTexts"`
}
