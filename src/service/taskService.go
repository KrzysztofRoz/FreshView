package service

import (
	"errors"
	"krzysztofRoz/FreshView/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TaskService struct {
	logger *zap.Logger
	db     *gorm.DB
}

type NewTaskInput struct {
	TaskName     string `json:"taskName" binding:"required"`
	TaskCategory string `json:"taskCategory" binding:"required"`
}

func NewTaskService(logger zap.Logger, db *gorm.DB) *TaskService {

	return &TaskService{
		logger: &logger,
		db:     db,
	}
}

func (ts TaskService) CreateNewTask(containerName string, input NewTaskInput) (model.DutieTask, error) {
	task := model.DutieTask{}
	container := model.DutieContainer{}

	ts.logger.Info("Query database for container",
		zap.String("containerName", containerName))

	result := ts.db.Where("container_name = ?", containerName).First(&container)

	if container.ID == 0 {
		ts.logger.Error("No such container in database",
			zap.String("containerName", containerName),
			zap.Error(result.Error))
		return task, ErrNoRecords
	}
	if result.Error != nil {
		ts.logger.Error("fail to retreive container",
			zap.String("containerName", containerName),
			zap.Error(result.Error))
		return task, result.Error
	}
	task.Category = input.TaskCategory
	task.TaskName = input.TaskName
	task.DutieContainerID = container.ID

	ts.logger.Info("Task created",
		zap.String("containerName", containerName),
		zap.String("taskName", task.TaskName),
		zap.String("taskCategory", task.Category))

	return task, nil
}

func (ts TaskService) AddTaskToDB(task model.DutieTask) error {
	ts.logger.Info("Inserting task to database",
		zap.String("taskName", task.TaskName))
	result := ts.db.Create(&task)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		ts.logger.Error("Task already exist in database",
			zap.String("taskName", task.TaskName),
			zap.Error(result.Error))
		return result.Error
	}
	if result.Error != nil {
		ts.logger.Error("Error in inserting task",
			zap.String("taskName", task.TaskName),
			zap.Error(result.Error))
		return result.Error
	}
	ts.logger.Info("Task sucesfully inserted in database",
		zap.String("taskName", task.TaskName))

	return nil

}
