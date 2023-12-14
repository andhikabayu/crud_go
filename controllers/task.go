package controllers

import (
	"crud_go/helpers"
	"crud_go/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	helpers.SuccessJSON(c, "Task retrieved successfully.", tasks)
}

// Create new task
func CreateTask(c *gin.Context) {
	// validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorJSON(c, err.Error())
		return
	}
	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	// create task
	task := models.Task{
		AssignedTo: input.AssignedTo,
		Task:       input.Task,
		Deadline:   deadline,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	helpers.SuccessJSON(c, "Task created successfully.", task)
}

// Find a task
func FindTask(c *gin.Context) {
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		helpers.ErrorJSON(c, "Record not found!")
		return
	}
	helpers.SuccessJSON(c, "Task retrived successfully", task)
}

// Update a task
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		helpers.ErrorJSON(c, "Data not found!")
		return
	}

	// validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorJSON(c, err.Error())
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Task
	updatedInput.Deadline = deadline
	updatedInput.AssignedTo = input.AssignedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	helpers.SuccessJSON(c, "Task updated successfully.", task)
}

// Delete a task
func DeleteTask(c *gin.Context) {
	// get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		helpers.ErrorJSON(c, "Data not found")
		return
	}

	db.Delete(&task)

	helpers.SuccessJSON(c, "Task deleted successfully.", true)
}
