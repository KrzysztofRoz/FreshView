package service

import (
	"errors"
	"krzysztofRoz/FreshView/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContainerService struct {
	logger *zap.Logger
	db     *gorm.DB
}

var ErrNoRecords = errors.New("no records found in database")

func NewContainerService(logger zap.Logger, db *gorm.DB) *ContainerService {

	return &ContainerService{
		logger: &logger,
		db:     db,
	}
}

func (cs ContainerService) CreateNewContainer(containerName string) model.DutieContainer {
	container := model.DutieContainer{ContainerName: containerName}
	cs.logger.Info("Sucesfully create container",
		zap.String("containerName", containerName))
	return container
}

func (cs ContainerService) SaveContainerToDB(container *model.DutieContainer) error {
	result := cs.db.Create(&container)
	if result.Error != nil {
		cs.logger.Error("Error in saving container to database",
			zap.String("containerName", container.ContainerName),
			zap.Error(result.Error))
		return result.Error
	}
	cs.logger.Info("Save container in database",
		zap.String("containerName", container.ContainerName))
	return nil
}

func (cs ContainerService) GetAllContainerNames() ([]string, error) {
	defer cs.logger.Sync()
	containerNames := make([]string, 0)
	var containers []model.DutieContainer
	cs.logger.Info("Query the database")
	result := cs.db.Select("container_name").Find(&containers)

	if result.Error != nil {
		cs.logger.Error("Error in retiving containers from database",
			zap.Error(result.Error))
		return containerNames, result.Error
	}
	if result.RowsAffected == 0 {
		cs.logger.Error("No containers in database",
			zap.Error(result.Error))
		return containerNames, ErrNoRecords
	}
	rows, err := result.Rows()
	if err != nil {
		cs.logger.Error("Error in retiving rows from database",
			zap.Error(result.Error))
		return containerNames, result.Error
	}
	cs.logger.Info("Succesfully query the database")

	defer rows.Close()
	for rows.Next() {
		var containerName string
		err = rows.Scan(&containerName)

		if err != nil {
			cs.logger.Error("Error when retiving row from database",
				zap.Error(result.Error))
			return containerNames, result.Error
		}
		cs.logger.Info("Append container to the slice",
			zap.String("containerName", containerName))
		containerNames = append(containerNames, containerName)
	}
	cs.logger.Info("Retreive containers from database",
		zap.Strings("containers", containerNames))

	return containerNames, nil

}
