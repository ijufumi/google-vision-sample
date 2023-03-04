package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/services"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"net/http"
	"os"
)

type DetectTextHandler interface {
	Gets(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewDetectTextHandler(service services.DetectTextService) DetectTextHandler {
	return &detectTextHandler{
		service: service,
	}
}

type detectTextHandler struct {
	service services.DetectTextService
}

func (h *detectTextHandler) Gets(ctx *gin.Context) {
	results, err := h.service.GetResults()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (h *detectTextHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.service.GetResultByID(id)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *detectTextHandler) Create(ctx *gin.Context) {
	inputFile, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tempFile, err := utils.NewTempFileWithName(inputFile.Filename)
	if err != nil {
		fmt.Println(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
	}()
	err = ctx.SaveUploadedFile(inputFile, tempFile.Name())
	if err != nil {
		fmt.Println(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.service.DetectTexts(tempFile, inputFile.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, models.Status{Status: true})
}

func (h *detectTextHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.service.DeleteResult(id)
	if err != nil {
		fmt.Println(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, models.Status{Status: true})
}
