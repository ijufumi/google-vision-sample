package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/services"
)

type DetectTextHandler interface {
	Get(ctx *gin.Context)
	Post(ctx *gin.Context)
}

func NewDetectTextHandler(service *services.DetectTextService) DetectTextHandler {
	return &detectTextHandler{
		service: service,
	}
}

type detectTextHandler struct {
	service *services.DetectTextService
}

func (h *detectTextHandler) Get(ctx *gin.Context) {

}

func (h *detectTextHandler) Post(ctx *gin.Context) {

}
