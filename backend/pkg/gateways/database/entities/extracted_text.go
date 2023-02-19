package entities

type ExtractedText struct {
	BaseEntity
	ID                 string
	ExtractionResultID string
	Text               string
	Top                float64
	Bottom             float64
	Left               float64
	Right              float64
}
