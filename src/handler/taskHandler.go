package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskHandler struct {
	Handler Handler
}

func NewTaskHandlerlogger(logger zap.Logger) *TaskHandler {
	return &TaskHandler{
		Handler: *NewHandler(logger),
	}
}

func (th TaskHandler) AddNewTask(ctx *gin.Context) {
	containerName := ctx.Param("containername")
	taskName := ctx.Param("taskname")

	ctx.JSON(http.StatusCreated, gin.H{
		"message":       "Endpoint reached",
		"containerName": containerName,
		"taskName":      taskName,
	})

}

func (th TaskHandler) RetreiveAllTasks(ctx *gin.Context) {
	containerName := ctx.Param("containername")

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Endpoint reached",
		"containerName": containerName,
	})

}
func (th TaskHandler) RemoveTask(ctx *gin.Context) {
	containerName := ctx.Param("containername")
	taskName := ctx.Param("taskname")

	ctx.JSON(http.StatusNoContent, gin.H{
		"message":       "Endpoint reached",
		"containerName": containerName,
		"taskName":      taskName,
	})

}
