package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContainerHandler struct {
	Handler Handler
}

func NewContainerHandler(logger zap.Logger) *ContainerHandler {
	return &ContainerHandler{
		Handler: *NewHandler(logger),
	}
}

func (ch ContainerHandler) AddNewContainer(ctx *gin.Context) {
	containerName := ctx.Param("containername")
	ctx.JSON(http.StatusCreated, gin.H{
		"message":       "Endpoint reached",
		"containerName": containerName,
	})
}

func (ch ContainerHandler) RetreiveAllContainers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Endpoint reached",
	})
}

func (ch ContainerHandler) RemoveContainer(ctx *gin.Context) {
	containerName := ctx.Param("containername")
	ctx.JSON(http.StatusNoContent, gin.H{
		"message":       "Endpoint reached",
		"containerName": containerName,
	})
}
