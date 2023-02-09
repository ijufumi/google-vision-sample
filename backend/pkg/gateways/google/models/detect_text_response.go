package models

type DetectTextResponses struct {
	Responses []DetectTextResponse `json:"responses"`
}

type DetectTextResponse struct {
	TextAnnotations    []TextAnnotation   `json:"textAnnotations"`
	FullTextAnnotation FullTextAnnotation `json:"fullTextAnnotation"`
}

type TextAnnotation struct {
	Locale       string       `json:"locale"`
	Description  string       `json:"description"`
	BoundingPoly BoundingPoly `json:"boundingPoly"`
}

type BoundingPoly struct {
	Vertices []Vertices `json:"vertices"`
}

type Vertices struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type FullTextAnnotation struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	Paragraphs  []Paragraph `json:"paragraphs"`
}

type BoundingBox struct {
	Vertices []Vertices `json:"vertices"`
}

type Paragraph struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	Words       []Word      `json:"words"`
}

type Word struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	Symbols     []Symbol    `json:"symbols"`
}

type Symbol struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	Text        string      `json:"text"`
}
