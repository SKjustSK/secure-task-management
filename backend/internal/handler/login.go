package handler

import (
	"net/http"

	"github.com/SKjustSK/secure-task-management/backend/internal/database"
	"github.com/SKjustSK/secure-task-management/backend/internal/models"
	"github.com/SKjustSK/secure-task-management/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// LoginInput defines the expected JSON for logging in
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginUser authenticates a user and returns a JWT
func LoginUser(c *gin.Context) {
	var input LoginInput

	// Validate the incoming JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Email and password are required")
		return
	}

	// Find the user in the database
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Verify the password hash
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate the JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate session token")
		return
	}

	// Return the token inside the standard 'data' field
	utils.Success(c, http.StatusOK, "Login successful", gin.H{"token": token})
}
