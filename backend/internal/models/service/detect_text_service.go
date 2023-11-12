package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ijufumi/google-vision-sample/internal/common/utils"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/google/clients"
	google "github.com/ijufumi/google-vision-sample/internal/infrastructures/google/models/entities"
	"github.com/ijufumi/google-vision-sample/internal/models/entities"
	"github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gopkg.in/gographics/imagick.v2/imagick"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type DetectTextService interface {
	GetResults(ctx context.Context) ([]*entities.Job, error)
	GetResultByID(ctx context.Context, id string) (*entities.Job, error)
	GetSignedURL(ctx context.Context, key string) (*entities.SignedURL, error)
	DetectTexts(ctx context.Context, file *os.File, contentType string) error
	DeleteResult(ctx context.Context, id string) error
}

func NewDetectTextService(
	storageAPIClient clients.StorageAPIClient,
	visionAPIClient clients.VisionAPIClient,
	jobRepository repositories.JobRepository,
	extractedTextRepository repositories.ExtractedTextRepository,
	inputFileRepository repositories.InputFileRepository,
	outputFileRepository repositories.OutputFileRepository,
	imageConversionService ImageConversionService,
) DetectTextService {
	return &detectTextServiceImpl{
		storageAPIClient:        storageAPIClient,
		visionAPIClient:         visionAPIClient,
		jobRepository:           jobRepository,
		extractedTextRepository: extractedTextRepository,
		inputFileRepository:     inputFileRepository,
		outputFileRepository:    outputFileRepository,
		imageConversionService:  imageConversionService,
	}
}

func (s *detectTextServiceImpl) GetResults(ctx context.Context) ([]*entities.Job, error) {
	var results []*entities.Job
	err := s.Process(ctx, func(logger *zap.Logger) error {
		_results, err := s.jobRepository.GetAll(ctx)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		results = _results
		return nil
	})
	if err != nil {
		return nil, err
	}

	return results, err
}

func (s *detectTextServiceImpl) GetResultByID(ctx context.Context, id string) (*entities.Job, error) {
	var response *entities.Job
	err := s.Process(ctx, func(logger *zap.Logger) error {
		result, err := s.jobRepository.GetByID(ctx, id)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		response = result
		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, err
}

func (s *detectTextServiceImpl) GetSignedURL(ctx context.Context, key string) (*entities.SignedURL, error) {
	var response *entities.SignedURL
	signedURL, err := s.storageAPIClient.SignedURL(key)
	if err != nil {
		return nil, err
	}
	response = &entities.SignedURL{URL: signedURL}
	return response, err
}

func (s *detectTextServiceImpl) DetectTexts(ctx context.Context, file *os.File, contentType string) error {
	id := utils.NewULID()
	key := fmt.Sprintf("%s/original/%s", id, filepath.Base(file.Name()))

	err := s.storageAPIClient.UploadFile(key, file, enums.ConvertToContentType(contentType))
	if err != nil {
		return err
	}

	splitFileName := strings.Split(file.Name(), "/")
	fileName := splitFileName[len(splitFileName)-1]

	job := &entities.Job{
		ID:              id,
		Name:            fileName,
		OriginalFileKey: key,
		Status:          enums.JobStatus_Runing,
	}
	err = s.jobRepository.Create(ctx, job)
	if err != nil {
		return err
	}

	tempFileForWork, err := utils.NewTempFileWithName(fileName)
	if err != nil {
		return err
	}
	err = utils.Copy(file, tempFileForWork)
	if err != nil {
		return err
	}

	go func() {
		_ = s.Process(ctx, func(logger *zap.Logger) error {
			job, err := s.jobRepository.GetByID(ctx, id)
			if err != nil {
				logger.Error(fmt.Sprintf("%v was occurred.", err))
			}
			err = s.processDetectText(ctx, logger, id, tempFileForWork)
			if err != nil {
				logger.Error(fmt.Sprintf("%v was occurred.", err))
				job.Status = enums.JobStatus_Failed
			} else {
				job.Status = enums.JobStatus_Succeeded
			}
			_ = s.jobRepository.Update(ctx, job)
			_ = os.Remove(tempFileForWork.Name())
			return nil
		})
	}()
	return nil
}

func (s *detectTextServiceImpl) processDetectText(ctx context.Context, logger *zap.Logger, id string, inputFile *os.File) error {
	contentType := s.imageConversionService.DetectContentType(inputFile.Name())
	// s.logger.Info(fmt.Sprintf("Content-Type is %s", contentType))
	switch {
	case contentType == enums.ContentType_Pdf:
		imageFiles, err := s.imageConversionService.ConvertPdfToImages(inputFile.Name())
		if err != nil {
			return err
		}
		errs := make([]error, 0)
		for idx, inputFile := range imageFiles {
			if inputFile == nil {
				continue
			}
			err := s.processDetectTextFromImage(ctx, logger, id, enums.ContentType_Png, inputFile, uint(idx+1))
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) != 0 {
			return errors.Join(errs...)
		}
		return nil
	case enums.IsImage(contentType):
		return s.processDetectTextFromImage(ctx, logger, id, contentType, inputFile, 1)
	}
	return errors.New(fmt.Sprintf("unsupported content-type : %s", contentType))
}

func (s *detectTextServiceImpl) processDetectTextFromImage(ctx context.Context, logger *zap.Logger, jobID string, contentType enums.ContentType, file *os.File, pageNo uint) error {
	inputFile, err := s.createInputFile(ctx, logger, jobID, contentType, file, pageNo)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = s.Transaction(ctx, func(ctx2 context.Context) error {
		outputFileKey, outputFile, err := s.detectText(logger, inputFile.FileKey)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		fileStat, _ := outputFile.Stat()
		splitOutputFileKey := strings.Split(outputFileKey, "/")
		outputFileID := utils.NewULID()
		err = s.outputFileRepository.Create(ctx2, &entities.OutputFile{
			ID:          outputFileID,
			JobID:       jobID,
			InputFileID: inputFile.ID,
			FileKey:     outputFileKey,
			FileName:    splitOutputFileKey[len(splitOutputFileKey)-1],
			ContentType: "application/json",
			Size:        uint(fileStat.Size()),
			Width:       0,
			Height:      0,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		orientation, err := s.imageConversionService.DetectOrientation(file.Name())
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		extractedTexts, err := s.convertToExtractedTexts(logger, jobID, inputFile, outputFileID, outputFile, orientation)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		return s.extractedTextRepository.Create(ctx2, extractedTexts...)
	})

	status := enums.InputFileStatus_Succeeded
	if err != nil {
		logger.Error(err.Error())
		status = enums.InputFileStatus_Failed
	}
	inputFile.Status = status
	_ = s.inputFileRepository.Update(ctx, inputFile)

	return err
}

func (s *detectTextServiceImpl) createInputFile(ctx context.Context, logger *zap.Logger, jobID string, contentType enums.ContentType, file *os.File, pageNo uint) (*entities.InputFile, error) {
	width, height, err := s.imageConversionService.DetectSize(file.Name())
	if err != nil {
		return nil, err
	}
	inputFileID := utils.NewULID()

	fileName := filepath.Base(file.Name())
	key := fmt.Sprintf("%s/%s/%s", jobID, inputFileID, filepath.Base(file.Name()))

	err = s.storageAPIClient.UploadFile(key, file, contentType)
	if err != nil {
		return nil, err
	}

	fileInfo, _ := file.Stat()
	err = s.inputFileRepository.Create(ctx, &entities.InputFile{
		ID:          inputFileID,
		JobID:       jobID,
		PageNo:      pageNo,
		FileKey:     key,
		FileName:    fileName,
		ContentType: contentType,
		Size:        uint(fileInfo.Size()),
		Width:       width,
		Height:      height,
		Status:      enums.InputFileStatus_Runing,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	inputFile, err := s.inputFileRepository.GetByID(ctx, inputFileID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return inputFile, nil
}

func (s *detectTextServiceImpl) detectText(logger *zap.Logger, inputFileKey string) (string, *os.File, error) {
	outputKey, err := s.visionAPIClient.DetectText(logger, inputFileKey)
	if err != nil {
		logger.Error(err.Error())
		return "", nil, err
	}

	queryFiles, err := s.storageAPIClient.QueryFiles(outputKey)
	if err != nil {
		logger.Error(err.Error())
		return "", nil, err
	}

	if len(queryFiles) == 0 {
		err = errors.New("output does not exist")
		logger.Error(err.Error())
		return "", nil, err
	}

	outputFile, err := s.storageAPIClient.DownloadFile(queryFiles[0])
	if err != nil {
		logger.Error(err.Error())
		return "", nil, err
	}
	return queryFiles[0], outputFile, nil
}

func (s *detectTextServiceImpl) convertToExtractedTexts(logger *zap.Logger, jobID string, inputFile *entities.InputFile, outputFileID string, outputFile *os.File, orientation imagick.OrientationType) (entities.ExtractedTexts, error) {
	bytes, err := io.ReadAll(outputFile)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var detectResponse google.DetectTextResponses
	if err = json.Unmarshal(bytes, &detectResponse); err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if len(detectResponse.Responses) == 0 {
		err = errors.New("response notfound")
		logger.Error(err.Error())
		return nil, err
	}
	response := detectResponse.Responses[0]
	if response.Error != nil {
		err = errors.New(response.Error.String())
		return nil, err
	}

	extractedTexts := make([]*entities.ExtractedText, 0)
	for _, page := range response.FullTextAnnotation.Pages {
		for _, block := range page.Blocks {
			for _, paragraph := range block.Paragraphs {
				texts := ""
				for _, word := range paragraph.Words {
					for _, symbol := range word.Symbols {
						texts += symbol.Text
					}
				}
				points := paragraph.BoundingBox.Vertices.ToDecimal()
				points = s.imageConversionService.ConvertPoints(points, orientation, inputFile.Width, inputFile.Height)

				xArray := make([]decimal.Decimal, 0)
				yArray := make([]decimal.Decimal, 0)
				for _, point := range points {
					xArray = append(xArray, point[0])
					yArray = append(yArray, point[1])
				}
				bottom, top := utils.MaxMinInDecimalArray(yArray...)
				right, left := utils.MaxMinInDecimalArray(xArray...)
				extractedText := &entities.ExtractedText{
					ID:           utils.NewULID(),
					JobID:        jobID,
					InputFileID:  inputFile.ID,
					OutputFileID: outputFileID,
					Text:         texts,
					Top:          top,
					Bottom:       bottom,
					Left:         left,
					Right:        right,
				}

				extractedTexts = append(extractedTexts, extractedText)
			}
		}
	}
	return extractedTexts, nil
}

func (s *detectTextServiceImpl) DeleteResult(ctx context.Context, id string) error {
	return s.Transaction(ctx, func(ctx2 context.Context) error {
		return s.Process(ctx2, func(logger *zap.Logger) error {
			outputFiles, err := s.outputFileRepository.GetByJobID(ctx2, id)
			if err != nil {
				return err
			}
			for _, file := range outputFiles {
				err = s.storageAPIClient.DeleteFile(file.FileKey)
				if err != nil {
					logger.Error(err.Error())
				}
				err := s.extractedTextRepository.DeleteByOutputFileID(ctx, file.ID)
				if err != nil {
					return err
				}
			}
			err = s.outputFileRepository.DeleteByJobID(ctx2, id)
			if err != nil {
				return err
			}
			inputFiles, err := s.inputFileRepository.GetByJobID(ctx2, id)
			if err != nil {
				return err
			}
			for _, file := range inputFiles {
				err = s.storageAPIClient.DeleteFile(file.FileKey)
				if err != nil {
					logger.Error(err.Error())
				}
			}
			err = s.inputFileRepository.DeleteByJobID(ctx2, id)
			if err != nil {
				// logger.Error(err.Error())
				return err
			}

			job, err := s.jobRepository.GetByID(ctx2, id)
			if err != nil {
				// logger.Error(err.Error())
				return err
			}
			err = s.storageAPIClient.DeleteFile(job.OriginalFileKey)
			if err != nil {
				logger.Error(err.Error())
			}
			return s.jobRepository.Delete(ctx2, id)
		})
	})
}

type detectTextServiceImpl struct {
	baseService
	storageAPIClient        clients.StorageAPIClient
	visionAPIClient         clients.VisionAPIClient
	jobRepository           repositories.JobRepository
	extractedTextRepository repositories.ExtractedTextRepository
	inputFileRepository     repositories.InputFileRepository
	outputFileRepository    repositories.OutputFileRepository
	imageConversionService  ImageConversionService
}
