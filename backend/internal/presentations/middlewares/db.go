package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/common/context"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/db"
)

func DB(config *configs.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l := context.GetLoggerFromGinContext(ctx)
		newDB := db.NewDB(config, l)
		context.SetDBToGinContext(ctx, newDB)
	}
}
