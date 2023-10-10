package service

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	"gopkg.in/gographics/imagick.v2/imagick"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ImageConversionService interface {
	DetectOrientation(filePath string) (imagick.OrientationType, error)
	DetectSize(filePath string) (width, height uint, err error)
	DetectContentType(filePath string) enums.ContentType
	ConvertPoints(points [][]float64, orientation imagick.OrientationType, width, height uint) [][]float64
	ConvertPdfToImages(pdfFilePath string) ([]*os.File, error)
}

type imageConversionService struct {
}

func NewImageConversionService() ImageConversionService {
	return &imageConversionService{}
}

func (s *imageConversionService) DetectOrientation(filePath string) (imagick.OrientationType, error) {
	imagick.Initialize()
	defer imagick.Terminate()

	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	err := magickWand.ReadImage(filePath)
	if err != nil {
		// s.logger.Error(fmt.Sprintf("error: %v", err))
		return imagick.ORIENTATION_UNDEFINED, err
	}
	orientation := magickWand.GetImageOrientation()
	// s.logger.Info(fmt.Sprintf("orientation is %v", orientation))
	return orientation, nil
}

func (s *imageConversionService) DetectSize(filePath string) (width, height uint, err error) {
	imagick.Initialize()
	defer imagick.Terminate()

	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	err = magickWand.ReadImage(filePath)
	if err != nil {
		// s.logger.Error(fmt.Sprintf("error: %v", err))
		return
	}
	height = magickWand.GetImageHeight()
	width = magickWand.GetImageWidth()

	return
}

func (s *imageConversionService) DetectContentType(filePath string) enums.ContentType {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		// s.logger.Error(fmt.Sprintf("failed detecting content-type: %v", err))
		return enums.ContentType_OctetStream
	}
	return enums.ConvertToContentType(http.DetectContentType(bytes))

}

func (s *imageConversionService) ConvertPoints(points [][]float64, orientation imagick.OrientationType, width, height uint) [][]float64 {
	if len(points) != 4 {
		// s.logger.Warn("point is invalid")
		return points
	}
	floatWidth := float64(width)
	floatHeight := float64(height)
	p1 := []float64{floatWidth - points[0][0], floatHeight - points[0][1]}
	p2 := []float64{floatWidth - points[1][0], floatHeight - points[1][1]}
	p3 := []float64{floatWidth - points[2][0], floatHeight - points[2][1]}
	p4 := []float64{floatWidth - points[3][0], floatHeight - points[3][1]}
	m := []float64{floatWidth / 2, floatHeight / 2}
	afterM := m
	angle := float64(0)
	switch orientation {
	case imagick.ORIENTATION_BOTTOM_LEFT:
		angle = 90
		afterM = []float64{afterM[1], afterM[0]}
	case imagick.ORIENTATION_RIGHT_TOP:
		angle = 270
		afterM = []float64{afterM[1], afterM[0]}
	case imagick.ORIENTATION_LEFT_BOTTOM:
		angle = 180
	default:
		// nothing
		return points
	}

	sin, cos := s.convertToSinCos(angle)
	p1 = s.convertPoint(p1, sin, cos, m, afterM)
	p2 = s.convertPoint(p2, sin, cos, m, afterM)
	p3 = s.convertPoint(p3, sin, cos, m, afterM)
	p4 = s.convertPoint(p4, sin, cos, m, afterM)

	return [][]float64{p1, p2, p3, p4}
}

func (s *imageConversionService) ConvertPdfToImages(pdfFilePath string) ([]*os.File, error) {
	imagick.Initialize()
	defer imagick.Terminate()

	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	err := magickWand.ReadImage(pdfFilePath)
	if err != nil {
		// s.logger.Error(fmt.Sprintf("error: %v", err))
		return nil, err
	}

	err = magickWand.SetImageFormat("png")
	if err != nil {
		// s.logger.Error(fmt.Sprintf("error: %v", err))
		return nil, err
	}

	pageNo := magickWand.GetNumberImages()
	// s.logger.Info(fmt.Sprintf("The page number of %s is %d", pdfFilePath, pageNo))

	imageFiles := make([]*os.File, pageNo)
	originalFilename := filepath.Base(pdfFilePath)
	for i := 0; i < int(pageNo); i++ {
		if ret := magickWand.SetIteratorIndex(i); !ret {
			continue
		}
		filePaths := []string{filepath.Dir(pdfFilePath), fmt.Sprintf("%s-%d.png", originalFilename, i)}
		imageFilePath := strings.Join(filePaths, "/")
		if err := magickWand.WriteImage(imageFilePath); err != nil {
			// s.logger.Error(fmt.Sprintf("error: %v", err))
			continue
		}
		imageFile, err := os.Open(imageFilePath)
		if err != nil {
			/// s.logger.Error(fmt.Sprintf("error: %v", err))
			return nil, err
		}
		imageFiles[i] = imageFile
	}
	return imageFiles, nil
}

func (s *imageConversionService) convertPoint(point []float64, sin, cos float64, beforeMiddlePoint, afterMiddlePoint []float64) []float64 {
	x := point[0] - beforeMiddlePoint[0]
	y := point[1] - beforeMiddlePoint[1]
	adjustPoint := make([]float64, 2)
	adjustPoint[0] = x*cos - y*sin
	adjustPoint[1] = x*sin + y*cos

	return []float64{adjustPoint[0] + afterMiddlePoint[0], adjustPoint[1] + afterMiddlePoint[1]}
}

func (s *imageConversionService) convertToSinCos(angle float64) (sin float64, cos float64) {
	// s.logger.Debug(fmt.Sprintf("angle is %v,", angle))
	sin, cos = math.Sincos(angle * math.Pi / 180)
	// s.logger.Debug(fmt.Sprintf("sin is %v, cos is %v", sin, cos))
	return
}