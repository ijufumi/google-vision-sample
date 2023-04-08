package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	googleModels "github.com/ijufumi/google-vision-sample/pkg/gateways/google/models"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type DetectTextService interface {
	GetResults() ([]*models.Job, error)
	GetResultByID(id string) (*models.Job, error)
	GetSignedURL(key string) (*models.SignedURL, error)
	DetectTexts(file *os.File, contentType string) error
	DeleteResult(id string) error
}

func NewDetectTextService(
	storageAPIClient clients.StorageAPIClient,
	visionAPIClient clients.VisionAPIClient,
	jobRepository repositories.JobRepository,
	extractedTextRepository repositories.ExtractedTextRepository,
	inputFileRepository repositories.InputFileRepository,
	outputFileRepository repositories.OutputFileRepository,
	imageConversionService ImageConversionService,
	logger *zap.Logger,
	db *gorm.DB,
) DetectTextService {
	return &detectTextService{
		storageAPIClient:        storageAPIClient,
		visionAPIClient:         visionAPIClient,
		jobRepository:           jobRepository,
		extractedTextRepository: extractedTextRepository,
		inputFileRepository:     inputFileRepository,
		outputFileRepository:    outputFileRepository,
		imageConversionService:  imageConversionService,
		logger:                  logger,
		db:                      db,
	}
}

func (s *detectTextService) GetResults() ([]*models.Job, error) {
	results, err := s.jobRepository.GetAll(s.db)
	if err != nil {
		return nil, err
	}

	extractionResults := make([]*models.Job, 0)
	for _, result := range results {
		extractionResults = append(extractionResults, s.buildExtractionResultResponse(result))
	}
	return extractionResults, nil
}

func (s *detectTextService) GetResultByID(id string) (*models.Job, error) {
	result, err := s.jobRepository.GetByID(s.db, id)
	if err != nil {
		return nil, err
	}
	return s.buildExtractionResultResponse(result), nil
}

func (s *detectTextService) GetSignedURL(key string) (*models.SignedURL, error) {
	signedURL, err := s.storageAPIClient.SignedURL(key)
	if err != nil {
		return nil, err
	}
	return &models.SignedURL{URL: signedURL}, nil
}

func (s *detectTextService) DetectTexts(file *os.File, contentType string) error {
	id := utils.NewULID()
	key := fmt.Sprintf("%s/original/%s", id, filepath.Base(file.Name()))

	err := s.storageAPIClient.UploadFile(key, file, contentType)
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
	err = s.jobRepository.Create(s.db, job)
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
		job, err := s.jobRepository.GetByID(s.db, id)
		if err != nil {
			s.logger.Error(fmt.Sprintf("%v was occurred.", err))
		}
		err = s.processDetectText(id, tempFileForWork)
		if err != nil {
			s.logger.Error(fmt.Sprintf("%v was occurred.", err))
			job.Status = enums.JobStatus_Failed
		} else {
			job.Status = enums.JobStatus_Succeeded
		}
		_ = s.jobRepository.Update(s.db, job)
		_ = os.Remove(tempFileForWork.Name())
	}()
	return nil
}

func (s *detectTextService) processDetectText(id string, inputFile *os.File) error {
	contentType := s.detectContentType(inputFile)
	s.logger.Info(fmt.Sprintf("Content-Type is %s", contentType))
	switch {
	case contentType == "application/pdf":
		imageFiles, err := s.imageConversionService.ConvertPdfToImages(inputFile.Name())
		if err != nil {
			return err
		}
		errs := make([]error, 0)
		for idx, inputFile := range imageFiles {
			if inputFile == nil {
				continue
			}
			err := s.processDetectTextFromImage(id, contentType, inputFile, uint(idx+1))
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) != 0 {
			return errors.Join(errs...)
		}
		return nil
	case strings.HasPrefix(contentType, "image/"):
		return s.processDetectTextFromImage(id, contentType, inputFile, 1)
	}
	return errors.New(fmt.Sprintf("unsupported content-type : %s", contentType))
}

func (s *detectTextService) processDetectTextFromImage(jobID string, contentType string, file *os.File, pageNo uint) error {
	width, height, err := s.imageConversionService.DetectSize(file.Name())
	if err != nil {
		return err
	}
	inputFileID := utils.NewULID()

	fileName := filepath.Base(file.Name())
	key := fmt.Sprintf("%s/%s/%s", jobID, inputFileID, filepath.Base(file.Name()))

	err = s.storageAPIClient.UploadFile(key, file, contentType)
	if err != nil {
		return err
	}

	fileInfo, _ := file.Stat()
	err = s.inputFileRepository.Create(s.db, &entities.InputFile{
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
		return err
	}
	inputFile, err := s.inputFileRepository.GetByID(s.db, inputFileID)
	if err != nil {
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		outputKey, err := s.visionAPIClient.DetectText(key)
		if err != nil {
			return err
		}

		queryFiles, err := s.storageAPIClient.QueryFiles(outputKey)
		if err != nil {
			return err
		}

		if len(queryFiles) == 0 {
			return errors.New("output does not exist")
		}

		outputFile, err := s.storageAPIClient.DownloadFile(queryFiles[0])
		if err != nil {
			return err
		}

		fileStat, _ := outputFile.Stat()
		outputFileKey := queryFiles[0]
		splitOutputFileKey := strings.Split(outputFileKey, "/")
		outputFileID := utils.NewULID()
		err = s.outputFileRepository.Create(tx, &entities.OutputFile{
			ID:          outputFileID,
			JobID:       jobID,
			InputFileID: inputFileID,
			FileKey:     outputFileKey,
			FileName:    splitOutputFileKey[len(splitOutputFileKey)-1],
			ContentType: "application/json",
			Size:        uint(fileStat.Size()),
			Width:       0,
			Height:      0,
		})
		if err != nil {
			return err
		}

		bytes, err := io.ReadAll(outputFile)
		if err != nil {
			return err
		}

		var detectResponse googleModels.DetectTextResponses
		if err = json.Unmarshal(bytes, &detectResponse); err != nil {
			return err
		}
		if len(detectResponse.Responses) == 0 {
			return errors.New("response notfound")
		}
		response := detectResponse.Responses[0]
		if response.Error != nil {
			return errors.New(response.Error.String())
		}

		extractedTexts := make([]*entities.ExtractedText, 0)
		orientation, err := s.imageConversionService.DetectOrientation(file.Name())
		if err != nil {
			return err
		}
		for _, page := range response.FullTextAnnotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					texts := ""
					for _, word := range paragraph.Words {
						for _, symbol := range word.Symbols {
							texts += symbol.Text
						}
					}
					points := paragraph.BoundingBox.Vertices.ToFloat()
					points = s.imageConversionService.ConvertPoints(points, orientation, width, height)

					xArray := make([]float64, 0)
					yArray := make([]float64, 0)
					for _, point := range points {
						xArray = append(xArray, point[0])
						yArray = append(yArray, point[1])
					}
					bottom, top := utils.MaxMinInArray(yArray...)
					right, left := utils.MaxMinInArray(xArray...)
					extractedText := &entities.ExtractedText{
						ID:           utils.NewULID(),
						JobID:        jobID,
						InputFileID:  inputFileID,
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
		return s.extractedTextRepository.Create(tx, extractedTexts...)
	})
	status := enums.InputFileStatus_Succeeded
	if err != nil {
		status = enums.InputFileStatus_Failed
	}
	inputFile.Status = status
	_ = s.inputFileRepository.Update(s.db, inputFile)
	return err
}

func (s *detectTextService) DeleteResult(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		outputFiles, err := s.outputFileRepository.GetByJobID(tx, id)
		if err != nil {
			return err
		}
		for _, file := range outputFiles {
			err = s.storageAPIClient.DeleteFile(file.FileKey)
			if err != nil {
				s.logger.Error(err.Error())
			}
			err := s.extractedTextRepository.DeleteByOutputFileID(tx, file.ID)
			if err != nil {
				return err
			}
		}
		err = s.outputFileRepository.DeleteByJobID(tx, id)
		if err != nil {
			return err
		}
		inputFiles, err := s.inputFileRepository.GetByJobID(tx, id)
		if err != nil {
			return err
		}
		for _, file := range inputFiles {
			err = s.storageAPIClient.DeleteFile(file.FileKey)
			if err != nil {
				s.logger.Error(err.Error())
			}
		}
		err = s.inputFileRepository.DeleteByJobID(tx, id)
		if err != nil {
			return err
		}

		job, err := s.jobRepository.GetByID(tx, id)
		if err != nil {
			return err
		}
		err = s.storageAPIClient.DeleteFile(job.OriginalFileKey)
		if err != nil {
			s.logger.Error(err.Error())
		}
		return s.jobRepository.Delete(tx, id)
	})
}

func (s *detectTextService) buildExtractionResultResponse(entity *entities.Job) *models.Job {
	inputFiles := make([]models.InputFile, 0)

	for _, inputFile := range entity.InputFiles {
		outputFiles := make([]*models.OutputFile, 0)
		for _, outputFile := range inputFile.OutputFiles {
			extractedTexts := make([]*models.ExtractedText, 0)
			for _, extractedText := range outputFile.ExtractedTexts {
				extractedTexts = append(extractedTexts, &models.ExtractedText{
					ID:           extractedText.ID,
					InputFileID:  extractedText.ID,
					OutputFileID: extractedText.ID,
					Text:         extractedText.Text,
					Top:          extractedText.Top,
					Bottom:       extractedText.Bottom,
					Left:         extractedText.Left,
					Right:        extractedText.Right,
					CreatedAt:    extractedText.CreatedAt.Unix(),
					UpdatedAt:    extractedText.UpdatedAt.Unix(),
				})
			}
			outputFiles = append(outputFiles, &models.OutputFile{
				ID:             outputFile.ID,
				JobID:          outputFile.JobID,
				InputFileID:    outputFile.InputFileID,
				FileKey:        outputFile.FileKey,
				FileName:       outputFile.FileName,
				ContentType:    outputFile.ContentType,
				Size:           outputFile.Size,
				CreatedAt:      outputFile.CreatedAt.Unix(),
				UpdatedAt:      outputFile.UpdatedAt.Unix(),
				ExtractedTexts: extractedTexts,
			})
		}
		inputFiles = append(inputFiles, models.InputFile{
			ID:          inputFile.ID,
			JobID:       inputFile.JobID,
			FileKey:     inputFile.FileKey,
			FileName:    inputFile.FileName,
			ContentType: inputFile.ContentType,
			Size:        inputFile.Size,
			CreatedAt:   inputFile.CreatedAt.Unix(),
			UpdatedAt:   inputFile.UpdatedAt.Unix(),
			OutputFiles: outputFiles,
		})
	}
	return &models.Job{
		ID:         entity.ID,
		Name:       entity.Name,
		Status:     entity.Status,
		CreatedAt:  entity.CreatedAt.Unix(),
		UpdatedAt:  entity.UpdatedAt.Unix(),
		InputFiles: inputFiles,
	}
}

func (s *detectTextService) detectContentType(file *os.File) string {
	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		s.logger.Error(fmt.Sprintf("failed detecting content-type: %v", err))
		return "application/octet-steam"
	}
	return http.DetectContentType(bytes)
}

type detectTextService struct {
	storageAPIClient        clients.StorageAPIClient
	visionAPIClient         clients.VisionAPIClient
	jobRepository           repositories.JobRepository
	extractedTextRepository repositories.ExtractedTextRepository
	inputFileRepository     repositories.InputFileRepository
	outputFileRepository    repositories.OutputFileRepository
	imageConversionService  ImageConversionService
	db                      *gorm.DB
	logger                  *zap.Logger
}
