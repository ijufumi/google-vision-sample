package services

import "os/exec"

type Orientation string

const (
	Orientation_Normal      = Orientation("0")
	Orientation_TopLeft     = Orientation("1")
	Orientation_TopRight    = Orientation("2")
	Orientation_BottomRight = Orientation("3")
	Orientation_BottomLeft  = Orientation("4")
	Orientation_LeftTop     = Orientation("5")
	Orientation_RightTop    = Orientation("6")
	Orientation_RightBottom = Orientation("7")
	Orientation_LeftBottom  = Orientation("8")
)

var identifyCommand = []string{"identify", "-format", "'%[orientation]'"}

var orientationMap = map[string]Orientation{
	"0": Orientation_Normal,
	"1": Orientation_TopLeft,
	"2": Orientation_TopRight,
	"3": Orientation_BottomRight,
	"4": Orientation_BottomLeft,
	"5": Orientation_LeftTop,
	"6": Orientation_RightTop,
	"7": Orientation_RightBottom,
	"8": Orientation_LeftBottom,
}

type ImageConversionService interface {
	DetectOrientation(filePath string) (Orientation, error)
}

type imageConversionService struct {
}

func NewImageConversionService() ImageConversionService {
	return &imageConversionService{}
}

func (s *imageConversionService) DetectOrientation(filePath string) (Orientation, error) {
	result, err := exec.Command(identifyCommand[0], identifyCommand[1:]...).Output()
	if err != nil {
		return Orientation_Normal, err
	}

	if orientation, ok := orientationMap[string(result)]; ok {
		return orientation, nil
	}

	return Orientation_Normal, nil
}
