package models

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"time"
)

type ExtractionResult struct {
	ID             string                       `json:"id"`
	Status         enums.ExtractionResultStatus `json:"status"`
	ImageUri       string                       `json:"imageUri"`
	OutputUri      string                       `json:"outputUri"`
	CreatedAt      time.Time                    `json:"createdAt"`
	UpdatedAt      time.Time                    `json:"updatedAt"`
	ExtractedTexts []ExtractedText              `json:"extractedTexts"`
}
