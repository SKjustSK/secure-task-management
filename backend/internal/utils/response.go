package utils

import (
	"github.com/gin-gonic/gin"
)

// APIResponse is the standard wrapper for all API responses
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success sends a standard success response
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error sends a standard error response
func Error(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error:   errorMessage,
	})
}
