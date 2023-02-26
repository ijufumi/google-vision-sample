package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/services"
	"net/http"
)

type SignedURLHandler interface {
	GetByKey(ctx *gin.Context)
}

type signedURLHandler struct {
	detectTextService services.DetectTextService
}

func NewSignedURL(detectTextService services.DetectTextService) SignedURLHandler {
	return &signedURLHandler{
		detectTextService: detectTextService,
	}
}

func (h *signedURLHandler) GetByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	signedURL, err := h.detectTextService.GetSignedURL(key)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, signedURL)
}
