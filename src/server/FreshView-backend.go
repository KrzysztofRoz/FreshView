package main

import (
	"krzysztofRoz/FreshView/handler"
	"krzysztofRoz/FreshView/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	repository.InitializeConfig()
	repository.ConnectDataBase()
	repository.SyncDB()
	containerHandler := handler.NewContainerHandler(*logger, repository.DB)
	taskHandler := handler.NewTaskHandlerlogger(*logger)
	router := gin.Default()
	authGroup := router.Group("/v1/api", repository.AuthMiddleweare())

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	//TODO add containers to database
	authGroup.POST("/add/container/:containername/", containerHandler.AddNewContainer)

	//TODO get all containers
	authGroup.GET("/retreive/containers/all", containerHandler.RetreiveAllContainers)

	//TODO remove container
	authGroup.DELETE("/remove/container/:containername", containerHandler.RemoveContainer)

	//TODO add task to container
	authGroup.POST("/add/task/:containername/:taskname", taskHandler.AddNewTask)

	//TODO get all task from container
	authGroup.GET("/retreive/tasks/:containername", taskHandler.RetreiveAllTasks)

	//TODO remove task from container
	authGroup.DELETE("/remove/task/:containername/:taskname", taskHandler.RemoveTask)

	// Run the server on port 8080
	router.Run(":8080")
}
