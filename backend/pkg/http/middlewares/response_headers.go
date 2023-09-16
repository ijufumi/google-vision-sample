package middlewares

import "github.com/gin-gonic/gin"

func ResponseHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseHeader := ctx.Writer.Header()
		responseHeader.Set("Cache-Control", "no-cache")
	}
}
