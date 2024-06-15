package model

import (
	"gorm.io/gorm"
)

type DutieTask struct {
	gorm.Model
	DutieContainerID uint
	TaskName         string `gorm:"index:idx__dutie_task_name,unique"`
	Category         string
}

type DutieContainer struct {
	gorm.Model
	ContainerName string `gorm:"uniqueIndex"`
	Duties        []DutieTask
}
