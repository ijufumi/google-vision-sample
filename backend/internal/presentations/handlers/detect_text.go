package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/utils"
	"github.com/ijufumi/google-vision-sample/internal/models/entity"
	"github.com/ijufumi/google-vision-sample/internal/models/service"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type DetectTextHandler interface {
	Gets(ginCtx *gin.Context)
	GetByID(ginCtx *gin.Context)
	Create(ginCtx *gin.Context)
	Delete(ginCtx *gin.Context)
}

func NewDetectTextHandler(service service.DetectTextService) DetectTextHandler {
	return &detectTextHandler{
		service: service,
	}
}

type detectTextHandler struct {
	baseHandler
	service service.DetectTextService
}

func (h *detectTextHandler) Gets(ginCtx *gin.Context) {
	_ = h.Process(ginCtx, func(ctx context.Context, logger *zap.Logger) error {
		results, err := h.service.GetResults(ctx, logger)
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusInternalServerError, err)
		}
		ginCtx.JSON(http.StatusOK, results)
		return nil
	})
}

func (h *detectTextHandler) GetByID(ginCtx *gin.Context) {
	_ = h.Process(ginCtx, func(ctx context.Context, logger *zap.Logger) error {
		id := ginCtx.Param("id")
		result, err := h.service.GetResultByID(ctx, logger, id)
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusInternalServerError, err)
		}
		ginCtx.JSON(http.StatusOK, result)

		return nil
	})
}

func (h *detectTextHandler) Create(ginCtx *gin.Context) {
	_ = h.Process(ginCtx, func(ctx context.Context, logger *zap.Logger) error {
		inputFile, err := ginCtx.FormFile("file")
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusBadRequest, err)
		}
		tempFile, err := utils.NewTempFileWithName(inputFile.Filename)
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusInternalServerError, err)
		}
		defer func() {
			err := os.Remove(tempFile.Name())
			if err != nil {
				logger.Error(err.Error())
			}
		}()
		err = ginCtx.SaveUploadedFile(inputFile, tempFile.Name())
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusInternalServerError, err)
		}
		err = h.service.DetectTexts(ctx, logger, tempFile, inputFile.Header.Get("Content-Type"))
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusBadRequest, err)
		}
		ginCtx.JSON(http.StatusOK, entity.Status{Status: true})
		return nil
	})
}

func (h *detectTextHandler) Delete(ginCtx *gin.Context) {
	_ = h.Process(ginCtx, func(ctx context.Context, logger *zap.Logger) error {
		id := ginCtx.Param("id")
		err := h.service.DeleteResult(ctx, logger, id)
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusInternalServerError, err)
		}
		ginCtx.JSON(http.StatusOK, entity.Status{Status: true})
		return nil
	})
}
