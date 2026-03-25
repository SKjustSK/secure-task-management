package handler

import (
	"net/http"

	"github.com/SKjustSK/secure-task-management/backend/internal/database"
	"github.com/SKjustSK/secure-task-management/backend/internal/models"
	"github.com/SKjustSK/secure-task-management/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// RegisterInput defines what JSON we expect from the client
type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterUser handles the creation of a new user
func RegisterUser(c *gin.Context) {
	var input RegisterInput

	// Validate the incoming JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Hash the password securely
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to process password")
		return
	}

	// Create the user model
	user := models.User{
		Email:    input.Email,
		Password: hashedPassword,
	}

	// Save to PostgreSQL via GORM
	if err := database.DB.Create(&user).Error; err != nil {
		utils.Error(c, http.StatusConflict, "Email is already registered")
		return
	}

	// Return success (filtering out the password for security)
	responseData := gin.H{
		"id":    user.ID,
		"email": user.Email,
	}
	utils.Success(c, http.StatusCreated, "User registered successfully", responseData)
}
