package entities

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type DetectTextResponses struct {
	Responses []DetectTextResponse `json:"responses"`
}

type DetectTextResponse struct {
	Error              *Error             `json:"error"`
	TextAnnotations    []TextAnnotation   `json:"textAnnotations"`
	FullTextAnnotation FullTextAnnotation `json:"fullTextAnnotation"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) String() string {
	return fmt.Sprintf("code: %d, messsage: %s", e.Code, e.Message)
}

type TextAnnotation struct {
	Locale       string       `json:"locale"`
	Description  string       `json:"description"`
	BoundingPoly BoundingPoly `json:"boundingPoly"`
}

type BoundingPoly struct {
	Vertices Vertices `json:"vertices"`
}

type Vertices []Vertex

func (v *Vertices) ToDecimal() [][]decimal.Decimal {
	value := make([][]decimal.Decimal, 0)

	for _, vertex := range *v {
		value = append(value, []decimal.Decimal{vertex.X, vertex.Y})
	}
	return value
}

type Vertex struct {
	X decimal.Decimal `json:"x"`
	Y decimal.Decimal `json:"y"`
}

type FullTextAnnotation struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	Width  decimal.Decimal `json:"width"`
	Height decimal.Decimal `json:"height"`
	Blocks []Block         `json:"blocks"`
}

type Block struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	Paragraphs  []Paragraph `json:"paragraphs"`
}

type BoundingBox struct {
	Vertices Vertices `json:"vertices"`
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
