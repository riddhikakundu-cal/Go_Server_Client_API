package main

import (
	"context"
	"log"
	"movie-api/go-server/routes"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "movie-api/go-server/docs" // Swagger docs import

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Movie API
// @version         1.0
// @description     A simple Movie Management API with async POST processing and polling support.
// @termsOfService  http://example.com/terms/
// @contact.name    Riddhika Kundu
// @contact.email   riddhika@example.com
// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT
// @host            localhost:8080
// @BasePath        /api
func main() {
	router := routes.SetupRouter()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  10 * time.Minute,
	}

	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

// package main

// import (
// 	"context"
// 	"log"
// 	"movie-api/go-server/routes"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"time"

// 	// Swagger docs (auto-generated)
// 	_ "movie-api/go-server/docs"
// )

// // @title           Movie API
// // @version         1.0
// // @description     A simple Movie Management API with async POST processing and polling support.
// // @host            localhost:8080
// // @BasePath        /api
// func main() {
// 	router := routes.SetupRouter() // All routes are handled inside

// 	server := &http.Server{
// 		Addr:         ":8080",
// 		Handler:      router,
// 		ReadTimeout:  2 * time.Minute,
// 		WriteTimeout: 5 * time.Minute,
// 		IdleTimeout:  10 * time.Minute,
// 	}

// 	go func() {
// 		log.Println("Server running on port 8080")
// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("Server error: %v", err)
// 		}
// 	}()

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit
// 	log.Println("Shutting down server...")

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	server.Shutdown(ctx)
// }
