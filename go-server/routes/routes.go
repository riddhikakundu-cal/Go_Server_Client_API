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
		api.GET("/movies/status/:taskId", controllers.CheckTaskStatus)
	}

	return router
}

// package routes

// import (
// 	"movie-api/go-server/controllers"
// 	_ "movie-api/go-server/docs" // Swagger docs (auto-generated)

// 	"github.com/gin-gonic/gin"
// 	swaggerFiles "github.com/swaggo/files"     // Swagger UI files
// 	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
// )

// func SetupRouter() *gin.Engine {
// 	router := gin.Default()

// 	api := router.Group("/api")
// 	{
// 		api.POST("/movies/batch", controllers.CreateMovieBatch)
// 		api.GET("/movies/status/:taskId", controllers.CheckTaskStatus)
// 	}

// 	// Swagger endpoint
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	return router
// }
