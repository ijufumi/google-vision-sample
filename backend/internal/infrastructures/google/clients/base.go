package clients

import (
	"context"
	contextManager "github.com/ijufumi/google-vision-sample/internal/common/context"
	"go.uber.org/zap"
)

type baseClient struct {
}

func (c *baseClient) Process(ctx context.Context, f func(logger *zap.Logger) error) error {
	logger := contextManager.GetLogger(ctx)
	return f(logger)
}
