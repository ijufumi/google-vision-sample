package services

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math"
	"os/exec"
	"strconv"
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

var identifyOrientationCommand = []string{"identify", "-format", "'%[orientation]'"}
var identifyWidthHeightCommand = []string{"identify", "-format", "'%[width],%[height]'"}

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
	DetectSize(filePath string) (width int64, height int64, err error)
	ConvertPoints(points []models.Vertices, orientation Orientation, width, height int64) [][]float64
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
	commands := append(identifyOrientationCommand, filePath)
	s.logger.Info(fmt.Sprintf("command: %s", strings.Join(commands, " ")))
	result, err := exec.Command(commands[0], commands[1:]...).Output()
	if err != nil {
		s.logger.Error(fmt.Sprintf("error: %v", err))
		return Orientation_None, err
	}

	resultStr := strings.ReplaceAll(string(result), "'", "")
	s.logger.Info(fmt.Sprintf("result: %s", resultStr))

	orientation := Orientation_None
	if _orientation, ok := orientationMap[resultStr]; ok {
		orientation = _orientation
	}

	s.logger.Info(fmt.Sprintf("orientation is %v", orientation))
	return orientation, nil
}

func (s *imageConversionService) DetectSize(filePath string) (width int64, height int64, err error) {
	commands := append(identifyWidthHeightCommand, filePath)
	s.logger.Info(fmt.Sprintf("command: %s", strings.Join(commands, " ")))
	result, err := exec.Command(commands[0], commands[1:]...).Output()
	if err != nil {
		s.logger.Error(fmt.Sprintf("error: %v", err))
		return 0, 0, err
	}

	resultStr := strings.ReplaceAll(string(result), "'", "")
	s.logger.Info(fmt.Sprintf("result: %s", resultStr))

	splitValue := strings.Split(resultStr, ",")
	if len(splitValue) != 2 {
		return 0, 0, errors.New(fmt.Sprintf("result was invalid value: %v", resultStr))
	}

	width, err = strconv.ParseInt(splitValue[0], 10, 64)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error: %v", err))
		return 0, 0, err
	}
	height, err = strconv.ParseInt(splitValue[1], 10, 64)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error: %v", err))
		return 0, 0, err
	}

	return
}

func (s *imageConversionService) ConvertPoints(points []models.Vertices, orientation Orientation, width, height int64) [][]float64 {
	if len(points) != 4 {
		s.logger.Warn("point is invalid")
		return [][]float64{}
	}
	p1 := []float64{points[0].X, points[0].Y}
	p2 := []float64{points[1].X, points[1].Y}
	p3 := []float64{points[2].X, points[2].Y}
	p4 := []float64{points[3].X, points[3].Y}
	m := []float64{float64(width / 2), float64(height / 2)}

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

	sin, cos := s.convertToSinCos(angle)
	p1 = s.convertPoint(p1, sin, cos, m)
	p2 = s.convertPoint(p2, sin, cos, m)
	p3 = s.convertPoint(p3, sin, cos, m)
	p4 = s.convertPoint(p4, sin, cos, m)

	return [][]float64{p1, p2, p3, p4}
}

func (s *imageConversionService) convertPoint(point []float64, sin, cos float64, middlePoint []float64) []float64 {
	x := point[0] - middlePoint[0]
	y := point[1] - middlePoint[1]
	adjustPoint := make([]float64, 2)
	adjustPoint[0] = x*cos - y*sin
	adjustPoint[1] = x*sin + y*cos

	return []float64{adjustPoint[0] + middlePoint[0], adjustPoint[1] + middlePoint[1]}
}

func (s *imageConversionService) convertToSinCos(angle float64) (sin float64, cos float64) {
	s.logger.Debug(fmt.Sprintf("angle is %v,", angle))
	sin, cos = math.Sincos(angle * math.Pi / 180)
	s.logger.Debug(fmt.Sprintf("sin is %v, cos is %v", sin, cos))
	return
}
