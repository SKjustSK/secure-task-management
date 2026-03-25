package handler

import (
	"net/http"

	"github.com/SKjustSK/secure-task-management/backend/internal/database"
	"github.com/SKjustSK/secure-task-management/backend/internal/models"
	"github.com/SKjustSK/secure-task-management/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// TaskInput defines the expected JSON for creating or updating a task
type TaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"` // e.g., pending, in_progress, completed
}

// CreateTask adds a new task for the authenticated user
func CreateTask(c *gin.Context) {
	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task data: "+err.Error())
		return
	}

	// Extract userID set by JWT middleware
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "User context not found")
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		UserID:      userID.(uint),
	}

	if err := database.DB.Create(&task).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to save task to database")
		return
	}

	utils.Success(c, http.StatusCreated, "Task created successfully", task)
}

// GetTasks retrieves all tasks belonging to the authenticated user
func GetTasks(c *gin.Context) {
	userID, _ := c.Get("userID")

	var tasks []models.Task
	if err := database.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to retrieve tasks")
		return
	}

	utils.Success(c, http.StatusOK, "Tasks retrieved successfully", tasks)
}

// GetTask retrieves a specific task by ID
func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	userID, _ := c.Get("userID")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Task not found or access denied")
		return
	}

	utils.Success(c, http.StatusOK, "Task retrieved", task)
}

// UpdateTask modifies an existing task
func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	userID, _ := c.Get("userID")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Task not found")
		return
	}

	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid update data")
		return
	}

	task.Title = input.Title
	task.Description = input.Description
	if input.Status != "" {
		task.Status = input.Status
	}

	if err := database.DB.Save(&task).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to update task")
		return
	}

	utils.Success(c, http.StatusOK, "Task updated successfully", task)
}

// DeleteTask removes a task (Soft Delete)
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	userID, _ := c.Get("userID")

	result := database.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{})

	if result.RowsAffected == 0 {
		utils.Error(c, http.StatusNotFound, "Task not found or already deleted")
		return
	}

	utils.Success(c, http.StatusOK, "Task deleted successfully", nil)
}
