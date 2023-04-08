package models

import "fmt"

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

func (v *Vertices) ToFloat() [][]float64 {
	value := make([][]float64, 0)

	for _, vertex := range *v {
		value = append(value, []float64{vertex.X, vertex.Y})
	}
	return value
}

type Vertex struct {
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
