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
	OutputFiles []*OutputFile
}

type InputFiles []*InputFile

func FromInputFileModel(inputFile *models.InputFile) *InputFile {
	return &InputFile{
		ID:          inputFile.ID,
		JobID:       inputFile.JobID,
		PageNo:      0, // fixme: set correct number
		FileKey:     inputFile.FileKey,
		FileName:    inputFile.FileName,
		ContentType: inputFile.ContentType,
		Size:        inputFile.Size,
		Width:       0,   // fixme: set correct number
		Height:      0,   // fixme: set correct number
		Status:      "",  // fixme: set correct status
		OutputFiles: nil, // fixme: set correct value
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
		OutputFiles: nil, // fixme: set correct value
	}
}

func (e *InputFiles) ToModel() models.InputFiles {
	var inputFiles models.InputFiles
	for _, inputFile := range *e {
		inputFiles = append(inputFiles, inputFile.ToModel())
	}
	return inputFiles
}
