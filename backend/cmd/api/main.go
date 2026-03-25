package main

import (
	"log"
	"os"

	"github.com/SKjustSK/secure-task-management/backend/internal/database"
	"github.com/SKjustSK/secure-task-management/backend/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	// Initialize Database Connection
	database.Connect()

	// Initialize Gin Router
	router := gin.Default()

	// Register Routes
	handler.RegisterRoutes(router)

	// Start the Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
