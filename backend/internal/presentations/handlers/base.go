package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	contextManager "github.com/ijufumi/google-vision-sample/internal/common/context"
	"go.uber.org/zap"
)

type baseHandler struct {
}

func (h *baseHandler) Process(ginCtx *gin.Context, process func(ctx context.Context, logger *zap.Logger) error) error {
	ctx := contextManager.GetContext(ginCtx)
	logger := contextManager.GetLogger(ctx)
	return process(ctx, logger)
}
