package handler

import (
	"errors"
	"krzysztofRoz/FreshView/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TaskHandler struct {
	handler     Handler
	taskService service.TaskService
}

func NewTaskHandler(logger zap.Logger, DB *gorm.DB) *TaskHandler {
	return &TaskHandler{
		handler:     *NewHandler(logger),
		taskService: *service.NewTaskService(logger, DB),
	}
}

func (th TaskHandler) AddNewTask(ctx *gin.Context) {
	containerName := ctx.Param("containername")
	var input service.NewTaskInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		th.handler.logger.Error("Invalid input")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := th.taskService.CreateNewTask(containerName, input)

	if err != nil {
		th.handler.logger.Error("Cannot find matching container",
			zap.String("containerName", containerName),
			zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":       "Cannot find matching container",
			"containerName": containerName,
			"error":         err.Error(),
		})
		return
	}

	err = th.taskService.AddTaskToDB(task)
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		th.handler.logger.Error("Duplicate task",
			zap.String("containerName", containerName),
			zap.String("taskName", input.TaskName),
			zap.Error(err))
		ctx.JSON(http.StatusConflict, gin.H{
			"message":       "Dupliacate task",
			"containerName": containerName,
			"taskName":      input.TaskName,
			"error":         err.Error(),
		})
		return
	}
	if err != nil {
		th.handler.logger.Error("Error inserting task",
			zap.String("containerName", containerName),
			zap.String("taskName", input.TaskName),
			zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":       "Dupliacate task",
			"containerName": containerName,
			"taskName":      input.TaskName,
			"error":         err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":       "Task sucesfully created",
		"containerName": containerName,
		"taskName":      task.TaskName,
		"taskCategory":  task.Category,
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
	err := th.taskService.DeleteTaskFormDB(containerName, taskName)

	if err != nil {
		th.handler.logger.Error("Error during delete event",
			zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Task sucesfully deleted",
		"containerName": containerName,
		"taskName":      taskName,
	})

}
