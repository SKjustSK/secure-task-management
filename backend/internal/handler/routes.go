package handler

import (
	"net/http"

	"github.com/SKjustSK/secure-task-management/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Global Middleware to avoid CORS error
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API is running smoothly"})
	})

	api := router.Group("/api")
	{
		// Auth Routes (Public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", RegisterUser)
			auth.POST("/login", LoginUser)
		}

		// Task Routes (Protected)
		tasks := api.Group("/tasks", middleware.RequireAuth())
		{
			tasks.GET("", GetTasks)
			tasks.POST("", CreateTask)
			tasks.GET("/:id", GetTask)
			tasks.PUT("/:id", UpdateTask)
			tasks.DELETE("/:id", DeleteTask)
		}
	}
}
