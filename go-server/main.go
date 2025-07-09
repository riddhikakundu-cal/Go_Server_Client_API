package main

import (
	"context"
	"log"
	"movie-api/go-server/routes"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routes.SetupRouter()

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
