package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/models/entity"
	"github.com/ijufumi/google-vision-sample/internal/services"
	"net/http"
)

type ConfigsHandler interface {
	SetupCORS(ctx *gin.Context)
}

type configsHandler struct {
	configurationService services.ConfigurationService
}

func NewConfigsHandler(configurationService services.ConfigurationService) ConfigsHandler {
	return &configsHandler{
		configurationService: configurationService,
	}
}

func (h *configsHandler) SetupCORS(ctx *gin.Context) {
	err := h.configurationService.SetupCORS()
	if err != nil {
		fmt.Println(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, entity.Status{Status: true})
}
