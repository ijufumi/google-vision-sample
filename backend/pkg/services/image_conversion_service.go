package services

type Orientation int

const (
	Orientation_TopLeft     = 1
	Orientation_TopRight    = 2
	Orientation_BottomRight = 3
	Orientation_BottomLeft  = 4
	Orientation_LeftTop     = 5
	Orientation_RightTop    = 6
	Orientation_RightBottom = 7
	Orientation_LeftBottom  = 8
)

type ImageConversionService interface {
}

type imageConversionService struct {
}

func NewImageConversionService() ImageConversionService {
	return &imageConversionService{}
}
