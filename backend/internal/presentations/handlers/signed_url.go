package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/models/service"
	"go.uber.org/zap"
	"net/http"
)

type SignedURLHandler interface {
	GetByKey(ginCtx *gin.Context)
}

type signedURLHandler struct {
	baseHandler
	detectTextService service.DetectTextService
}

func NewSignedURL(detectTextService service.DetectTextService) SignedURLHandler {
	return &signedURLHandler{
		detectTextService: detectTextService,
	}
}

func (h *signedURLHandler) GetByKey(ginCtx *gin.Context) {
	_ = h.Process(ginCtx, func(ctx context.Context, logger *zap.Logger) error {
		key := ginCtx.Query("key")
		signedURL, err := h.detectTextService.GetSignedURL(ctx, key)
		if err != nil {
			logger.Error(err.Error())
			return ginCtx.AbortWithError(http.StatusBadRequest, err)
		}
		ginCtx.JSON(http.StatusOK, signedURL)
		return nil
	})
}
