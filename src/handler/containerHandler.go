package handler

import (
	"errors"
	"krzysztofRoz/FreshView/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContainerHandler struct {
	handler          Handler
	containerService service.ContainerService
}

func NewContainerHandler(logger zap.Logger, DB *gorm.DB) *ContainerHandler {
	return &ContainerHandler{
		handler:          *NewHandler(logger),
		containerService: *service.NewContainerService(logger, DB),
	}
}

func (ch ContainerHandler) AddNewContainer(ctx *gin.Context) {
	containerName := ctx.Param("containername")

	container := ch.containerService.CreateNewContainer(containerName)
	result := ch.containerService.SaveContainerToDB(&container)

	if result != nil {
		ch.handler.logger.Error("Database insert error")
		ctx.JSON(http.StatusConflict, gin.H{
			"error":   "Error inserting to database",
			"message": result,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":       "Container sucesfully added",
		"containerName": containerName,
	})
}

func (ch ContainerHandler) RetreiveAllContainers(ctx *gin.Context) {
	defer ch.handler.logger.Sync()
	containers, err := ch.containerService.GetAllContainerNames()

	if err != nil && errors.Is(err, service.ErrNoRecords) {
		ch.handler.logger.Error("There is no containers in database",
			zap.Error(err))
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

	}
	if err != nil {
		ch.handler.logger.Error("Error during retreive event",
			zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":        "Conteiners sucesfully retreive",
		"conteinerNames": containers,
	})
}

func (ch ContainerHandler) RemoveContainer(ctx *gin.Context) {
	containerName := ctx.Param("containername")
	ctx.JSON(http.StatusNoContent, gin.H{
		"message":       "Endpoint reached",
		"containerName": containerName,
	})
}
