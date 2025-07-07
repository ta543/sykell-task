package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ta543/task/backend/handlers"
	"github.com/ta543/task/backend/models"
)

func main() {
	dsn := os.Getenv("WEBCRAWLER_DSN")
	if dsn == "" {
		// Changed from "mysql" (Docker service name) to "localhost"
		dsn = "root:1234@tcp(localhost:3306)/webcrawler?parseTime=true"
	}

	if err := models.InitDB(dsn); err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	r := gin.Default()
	r.Use(handlers.AuthMiddleware())

	r.POST("/api/urls", handlers.AddURL)
	r.GET("/api/urls", handlers.ListURLs)
	r.GET("/api/urls/:id", handlers.GetURL)
	r.DELETE("/api/urls/:id", handlers.DeleteURL)
	r.POST("/api/urls/:id/restart", handlers.RestartURL)
	r.POST("/api/urls/:id/stop", handlers.StopURL)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
