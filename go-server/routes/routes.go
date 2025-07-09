package routes

import (
	"movie-api/go-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/movies/batch", controllers.CreateMovieBatch)
		// api.GET("/status/:taskId", controllers.CheckTaskStatus)
		api.GET("/movies/status/:taskId", controllers.CheckTaskStatus)

	}
	return router
}
