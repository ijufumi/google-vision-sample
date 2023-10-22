package entities

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	"gorm.io/gorm"
)

type InputFile struct {
	gorm.Model
	ID          string
	JobID       string
	PageNo      uint
	FileKey     string
	FileName    string
	ContentType enums.ContentType
	Size        uint
	Width       uint
	Height      uint
	Status      enums.InputFileStatus
	OutputFiles OutputFiles
}

type InputFiles []*InputFile

func FromInputFilesModel(inputFilesModel models.InputFiles) InputFiles {
	var inputFiles InputFiles
	for _, inputFile := range inputFilesModel {
		inputFiles = append(inputFiles, FromInputFileModel(inputFile))
	}
	return inputFiles
}

func FromInputFileModel(inputFileModel *models.InputFile) *InputFile {
	return &InputFile{
		ID:          inputFileModel.ID,
		JobID:       inputFileModel.JobID,
		PageNo:      inputFileModel.PageNo,
		FileKey:     inputFileModel.FileKey,
		FileName:    inputFileModel.FileName,
		ContentType: inputFileModel.ContentType,
		Size:        inputFileModel.Size,
		Width:       inputFileModel.Width,
		Height:      inputFileModel.Height,
		Status:      inputFileModel.Status,
		OutputFiles: FromOutputFilesModel(inputFileModel.OutputFiles),
	}
}

func (e *InputFile) ToModel() *models.InputFile {
	return &models.InputFile{
		ID:          e.ID,
		JobID:       e.JobID,
		FileKey:     e.FileKey,
		FileName:    e.FileName,
		ContentType: e.ContentType,
		Size:        e.Size,
		CreatedAt:   e.CreatedAt.Unix(),
		UpdatedAt:   e.UpdatedAt.Unix(),
		OutputFiles: e.OutputFiles.ToModel(),
	}
}

func (e *InputFiles) ToModel() models.InputFiles {
	var inputFiles models.InputFiles
	for _, inputFile := range *e {
		inputFiles = append(inputFiles, inputFile.ToModel())
	}
	return inputFiles
}
