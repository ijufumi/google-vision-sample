package service

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	"github.com/shopspring/decimal"
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
	ConvertPoints(points [][]decimal.Decimal, orientation imagick.OrientationType, width, height uint) [][]decimal.Decimal
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

func (s *imageConversionService) ConvertPoints(points [][]decimal.Decimal, orientation imagick.OrientationType, width, height uint) [][]decimal.Decimal {
	if len(points) != 4 {
		// s.logger.Warn("point is invalid")
		return points
	}
	floatWidth := decimal.NewFromFloat(float64(width))
	floatHeight := decimal.NewFromFloat(float64(height))
	p1 := []decimal.Decimal{floatWidth.Sub(points[0][0]), floatHeight.Sub(points[0][1])}
	p2 := []decimal.Decimal{floatWidth.Sub(points[1][0]), floatHeight.Sub(points[1][1])}
	p3 := []decimal.Decimal{floatWidth.Sub(points[2][0]), floatHeight.Sub(points[2][1])}
	p4 := []decimal.Decimal{floatWidth.Sub(points[3][0]), floatHeight.Sub(points[3][1])}
	m := []decimal.Decimal{floatWidth.Div(decimal.NewFromFloat(float64(2))), floatHeight.Div(decimal.NewFromFloat(float64(2)))}
	afterM := m
	angle := float64(0)
	switch orientation {
	case imagick.ORIENTATION_BOTTOM_LEFT:
		angle = 90
		afterM = []decimal.Decimal{afterM[1], afterM[0]}
	case imagick.ORIENTATION_RIGHT_TOP:
		angle = 270
		afterM = []decimal.Decimal{afterM[1], afterM[0]}
	case imagick.ORIENTATION_LEFT_BOTTOM:
		angle = 180
	default:
		// nothing
		return points
	}

	sin, cos := s.convertToSinCos(decimal.NewFromFloat(angle))
	p1 = s.convertPoint(p1, sin, cos, m, afterM)
	p2 = s.convertPoint(p2, sin, cos, m, afterM)
	p3 = s.convertPoint(p3, sin, cos, m, afterM)
	p4 = s.convertPoint(p4, sin, cos, m, afterM)

	return [][]decimal.Decimal{p1, p2, p3, p4}
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

func (s *imageConversionService) convertPoint(point []decimal.Decimal, sin, cos decimal.Decimal, beforeMiddlePoint, afterMiddlePoint []decimal.Decimal) []decimal.Decimal {
	x := point[0].Sub(beforeMiddlePoint[0])
	y := point[1].Sub(beforeMiddlePoint[1])
	adjustPoint := make([]decimal.Decimal, 2)
	adjustPoint[0] = x.Mul(cos).Sub(y.Mul(sin))
	adjustPoint[1] = x.Mul(sin).Add(y.Mul(cos))

	return []decimal.Decimal{adjustPoint[0].Add(afterMiddlePoint[0]), adjustPoint[1].Add(afterMiddlePoint[1])}
}

func (s *imageConversionService) convertToSinCos(angle decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	// s.logger.Debug(fmt.Sprintf("angle is %v,", angle))
	sinFloat, cosFloat := math.Sincos(angle.Mul(decimal.NewFromFloat(math.Pi)).Div(decimal.New(180, 32)).InexactFloat64())
	// s.logger.Debug(fmt.Sprintf("sin is %v, cos is %v", sin, cos))
	return decimal.NewFromFloat(sinFloat), decimal.NewFromFloat(cosFloat)
}
