package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DutieTask struct {
	gorm.Model
	DutieContainerID uint
	ProductName      string
	Category         string
	CreatedAt        datatypes.Time
	UpdatedAt        datatypes.Time
}

type DutieContainer struct {
	gorm.Model
	ContainerName string `gorm:"uniqueIndex"`
	Duties        []DutieTask
}
