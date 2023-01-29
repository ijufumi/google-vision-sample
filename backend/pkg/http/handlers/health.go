package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler interface {
	Get(ctx *gin.Context)
}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

type healthHandler struct {
}
