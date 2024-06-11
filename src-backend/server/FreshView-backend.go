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
	taskHandler := handler.NewTaskHandler(*logger, repository.DB)
	router := gin.Default()
	authGroup := router.Group("/v1/api", repository.AuthMiddleweare(), repository.CORSMiddleware())
	corsGroup := router.Group("/v1/api", repository.CORSMiddleware())

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	//cors endpoints
	corsGroup.OPTIONS("/add/container/:containername/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "return headers",
		})
	})

	corsGroup.OPTIONS("/retreive/containers/all", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "return headers",
		})
	})

	corsGroup.OPTIONS("/retreive/container/:containername", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "return headers",
		})
	})

	corsGroup.OPTIONS("/remove/container/:containername", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "return headers",
		})
	})

	corsGroup.OPTIONS("/add/task/:containername", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "return headers",
		})
	})

	corsGroup.OPTIONS("/remove/task/:containername/:taskname", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "return headers",
		})
	})

	// Add containers to database
	authGroup.POST("/add/container/:containername/", containerHandler.AddNewContainer)

	// Get all containers
	authGroup.GET("/retreive/containers/all", containerHandler.RetreiveAllContainers)

	// Get specyfic container info with
	authGroup.GET("/retreive/container/:containername", containerHandler.RetreiveSingleContainer)

	// Remove container
	authGroup.DELETE("/remove/container/:containername", containerHandler.RemoveContainer)

	// Add task to container
	authGroup.POST("/add/task/:containername", taskHandler.AddNewTask)

	//TODO Get task details
	// authGroup.GET("/retreive/task/:containername/:taskname", taskHandler.RetreiveTaskDetailes)

	//TODO remove task from container
	authGroup.DELETE("/remove/task/:containername/:taskname", taskHandler.RemoveTask)

	// Run the server on port 8080
	router.Run(":8080")
}
