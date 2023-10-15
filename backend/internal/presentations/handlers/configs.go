package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/models/entities"
	"github.com/ijufumi/google-vision-sample/internal/models/service"
	"net/http"
)

type ConfigsHandler interface {
	SetupCORS(ctx *gin.Context)
}

type configsHandler struct {
	configurationService service.ConfigurationService
}

func NewConfigsHandler(configurationService service.ConfigurationService) ConfigsHandler {
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
	ctx.JSON(http.StatusOK, entities.Status{Status: true})
}
