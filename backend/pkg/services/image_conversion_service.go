package services

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/models"
	"go.uber.org/zap"
	"math"
	"os/exec"
	"strings"
)

type Orientation string

const (
	Orientation_None        = Orientation("None")
	Orientation_TopLeft     = Orientation("TopLeft")
	Orientation_TopRight    = Orientation("TopRight")
	Orientation_BottomRight = Orientation("BottomRight")
	Orientation_BottomLeft  = Orientation("BottomLeft")
	Orientation_LeftTop     = Orientation("LeftTop")
	Orientation_RightTop    = Orientation("RightTop")
	Orientation_RightBottom = Orientation("RightBottom")
	Orientation_LeftBottom  = Orientation("LeftBottom")
)

var identifyCommand = []string{"identify", "-format", "'%[orientation]'"}

var orientationMap = map[string]Orientation{
	"TopLeft":     Orientation_TopLeft,
	"TopRight":    Orientation_TopRight,
	"BottomRight": Orientation_BottomRight,
	"BottomLeft":  Orientation_BottomLeft,
	"LeftTop":     Orientation_LeftTop,
	"RightTop":    Orientation_RightTop,
	"RightBottom": Orientation_RightBottom,
	"LeftBottom":  Orientation_LeftBottom,
}

type ImageConversionService interface {
	DetectOrientation(filePath string) (Orientation, error)
	ConvertPoints(points []models.Vertices, orientation Orientation) [][]float64
}

type imageConversionService struct {
	logger *zap.Logger
}

func NewImageConversionService(logger *zap.Logger) ImageConversionService {
	return &imageConversionService{
		logger: logger,
	}
}

func (s *imageConversionService) DetectOrientation(filePath string) (Orientation, error) {
	commands := identifyCommand
	commands = append(commands, filePath)
	s.logger.Info(fmt.Sprintf("command: %s", strings.Join(commands, " ")))
	result, err := exec.Command(commands[0], commands[1:]...).Output()

	resultStr := strings.ReplaceAll(string(result), "'", "")
	s.logger.Info(fmt.Sprintf("result: %s", resultStr))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error: %v", err))
		return Orientation_None, err
	}

	orientation := Orientation_None
	if _orientation, ok := orientationMap[resultStr]; ok {
		orientation = _orientation
	}

	s.logger.Info(fmt.Sprintf("orientation is %v", orientation))
	return orientation, nil
}

func (s *imageConversionService) ConvertPoints(points []models.Vertices, orientation Orientation) [][]float64 {
	if len(points) != 4 {
		return [][]float64{}
	}
	p1 := []float64{points[0].X, points[0].Y}
	p2 := []float64{points[1].X, points[1].Y}
	p3 := []float64{points[2].X, points[2].Y}
	p4 := []float64{points[3].X, points[3].Y}
	m := []float64{(p3[0] - p1[0]) / 2, (p3[1] - p1[1]) / 2}

	angle := float64(0)
	switch orientation {
	case Orientation_BottomLeft:
		angle = 180
	case Orientation_RightTop:
		angle = 270
	case Orientation_LeftBottom:
		angle = 90
	default:
		// nothing
	}

	p1 = s.convertPoint(p1, angle, m)
	p2 = s.convertPoint(p2, angle, m)
	p3 = s.convertPoint(p3, angle, m)
	p4 = s.convertPoint(p4, angle, m)

	return [][]float64{p1, p2, p3, p4}
}

func (s *imageConversionService) convertPoint(point []float64, angle float64, middlePoint []float64) []float64 {
	sin, cos := math.Sincos(angle * math.Pi / 180)

	adjustPoint := []float64{point[0] - middlePoint[0], point[1] - middlePoint[1]}
	adjustPoint[0] = adjustPoint[0]*cos - adjustPoint[1]*sin
	adjustPoint[1] = adjustPoint[0]*sin - adjustPoint[1]*cos

	return []float64{adjustPoint[0] + middlePoint[0], adjustPoint[1] + middlePoint[1]}
}
