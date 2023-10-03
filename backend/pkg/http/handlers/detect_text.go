package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/services"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
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

func NewDetectTextHandler(service services.DetectTextService) DetectTextHandler {
	return &detectTextHandler{
		service: service,
	}
}

type detectTextHandler struct {
	baseHandler
	service services.DetectTextService
}

func (h *detectTextHandler) Gets(ginCtx *gin.Context) {
	_ = h.Process(ginCtx, func(ctx context.Context, logger *zap.Logger) error {
		results, err := h.service.GetResults(ctx)
		if err != nil {
			_ = ginCtx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}
		ginCtx.JSON(http.StatusOK, results)
		return nil
	})
}

func (h *detectTextHandler) GetByID(ginCtx *gin.Context) {
	//ctx := context.GetContextWithLogger(ginCtx)
	id := ginCtx.Param("id")
	result, err := h.service.GetResultByID(id)
	if err != nil {
		_ = ginCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ginCtx.JSON(http.StatusOK, result)
}

func (h *detectTextHandler) Create(ginCtx *gin.Context) {
	//ctx := context.GetContextWithLogger(ginCtx)
	inputFile, err := ginCtx.FormFile("file")
	if err != nil {
		fmt.Println(err)
		_ = ginCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tempFile, err := utils.NewTempFileWithName(inputFile.Filename)
	if err != nil {
		fmt.Println(err)
		_ = ginCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()
	err = ginCtx.SaveUploadedFile(inputFile, tempFile.Name())
	if err != nil {
		fmt.Println(err)
		_ = ginCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.service.DetectTexts(tempFile, inputFile.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println(err)
		_ = ginCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ginCtx.JSON(http.StatusOK, models.Status{Status: true})
}

func (h *detectTextHandler) Delete(ginCtx *gin.Context) {
	//ctx := context.GetContextWithLogger(ginCtx)
	id := ginCtx.Param("id")
	err := h.service.DeleteResult(id)
	if err != nil {
		fmt.Println(err)
		_ = ginCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ginCtx.JSON(http.StatusOK, models.Status{Status: true})
}
