package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/http/context"
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

func (h *signedURLHandler) GetByKey(ginCtx *gin.Context) {
	ctx := context.GetContextWithLogger(ginCtx)
	key := ginCtx.Query("key")
	signedURL, err := h.detectTextService.GetSignedURL(ctx, key)
	if err != nil {
		_ = ginCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ginCtx.JSON(http.StatusOK, signedURL)
}
