package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	"gorm.io/gorm"
)

type OutputFile struct {
	gorm.Model
	ID             string
	JobID          string
	InputFileID    string
	FileKey        string
	FileName       string
	ContentType    enums.ContentType
	Size           uint
	Width          uint
	Height         uint
	ExtractedTexts ExtractedTexts
}

type OutputFiles []*OutputFile

func (e *OutputFile) ToModel() *models.OutputFile {
	return &models.OutputFile{
		ID:             e.ID,
		JobID:          e.JobID,
		InputFileID:    e.InputFileID,
		FileKey:        e.FileKey,
		FileName:       e.FileName,
		ContentType:    e.ContentType,
		Size:           e.Size,
		CreatedAt:      e.CreatedAt.Unix(),
		UpdatedAt:      e.UpdatedAt.Unix(),
		ExtractedTexts: e.ExtractedTexts.ToModel(),
	}
}

func (e *OutputFiles) ToModel() models.OutputFiles {
	var outputFiles models.OutputFiles

	for _, outputFile := range *e {
		outputFiles = append(outputFiles, outputFile.ToModel())
	}
	return outputFiles
}

func FromOutputFilesModel(outputFilesModel models.OutputFiles) OutputFiles {
	var outputFiles OutputFiles
	for _, outputFile := range outputFilesModel {
		outputFiles = append(outputFiles, FromOutputFileModel(outputFile))
	}
	return outputFiles
}

func FromOutputFileModel(outputFileModel *models.OutputFile) *OutputFile {
	return &OutputFile{
		ID:             outputFileModel.ID,
		JobID:          outputFileModel.JobID,
		InputFileID:    outputFileModel.InputFileID,
		FileKey:        outputFileModel.FileKey,
		FileName:       outputFileModel.FileName,
		ContentType:    outputFileModel.ContentType,
		Size:           outputFileModel.Size,
		Width:          outputFileModel.Width,
		Height:         outputFileModel.Height,
		ExtractedTexts: FromExtractedTextsModel(outputFileModel.ExtractedTexts),
	}
}
