package entities

import (
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	"gorm.io/gorm"
)

type ExtractedText struct {
	gorm.Model
	ID           string
	JobID        string
	InputFileID  string
	OutputFileID string
	Text         string
	Top          float64
	Bottom       float64
	Left         float64
	Right        float64
}

type ExtractedTexts []*ExtractedText

func (e *ExtractedText) ToModel() *models.ExtractedText {
	return &models.ExtractedText{
		ID:           e.ID,
		InputFileID:  e.InputFileID,
		OutputFileID: e.OutputFileID,
		Text:         e.Text,
		Top:          e.Top,
		Bottom:       e.Bottom,
		Left:         e.Left,
		Right:        e.Right,
		CreatedAt:    e.CreatedAt.Unix(),
		UpdatedAt:    e.UpdatedAt.Unix(),
	}
}

func (e *ExtractedTexts) ToModel() models.ExtractedTexts {
	var extractedTexts models.ExtractedTexts
	for _, ExtractedText := range *e {
		extractedTexts = append(extractedTexts, ExtractedText.ToModel())
	}
	return extractedTexts
}
