package entities

import "github.com/shopspring/decimal"

type ExtractedText struct {
	ID           string          `json:"id"`
	JobID        string          `json:"jobID"`
	InputFileID  string          `json:"inputFileID"`
	OutputFileID string          `json:"outputFileID"`
	Text         string          `json:"text"`
	Top          decimal.Decimal `json:"top"`
	Bottom       decimal.Decimal `json:"bottom"`
	Left         decimal.Decimal `json:"left"`
	Right        decimal.Decimal `json:"right"`
	CreatedAt    int64           `json:"createdAt"`
	UpdatedAt    int64           `json:"updatedAt"`
}

type ExtractedTexts []*ExtractedText
