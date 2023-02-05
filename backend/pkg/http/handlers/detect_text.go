package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/services"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"net/http"
	"os"
)

type DetectTextHandler interface {
	Get(ctx *gin.Context)
	Post(ctx *gin.Context)
}

func NewDetectTextHandler(service services.DetectTextService) DetectTextHandler {
	return &detectTextHandler{
		service: service,
	}
}

type detectTextHandler struct {
	service services.DetectTextService
}

func (h *detectTextHandler) Get(ctx *gin.Context) {
	results, err := h.service.GetResults()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (h *detectTextHandler) Post(ctx *gin.Context) {
	inputFile, err := ctx.FormFile("file")
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tempFile, err := utils.NewTempFile()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()
	_ = ctx.SaveUploadedFile(inputFile, tempFile.Name())
	err = h.service.DetectTexts(tempFile)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}
