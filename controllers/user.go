package controllers

import (
	"crud_sederhana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Name   string `json:"name" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Level  string `json:"level" binding:"required"`
}

type UpdateUserInput struct {
	Name   string `json:"name" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Level  string `json:"level" binding:"required"`
}

// Get All User
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data []models.User
	db.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Store User
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create
	data := models.User{Name: input.Name, Gender: input.Gender, Level: input.Level}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Find User by Id
func FindUserById(c *gin.Context) {
	var data models.User

	db := c.MustGet("db").(*gorm.DB)
	// Get Data if Exist
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Update Users
func UpdateUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get Data if Exist
	var data models.User
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!!"})
		return
	}

	// Validate Input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.Name = input.Name
	updatedInput.Gender = input.Gender
	updatedInput.Level = input.Level

	db.Model(&data).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Delete User
func DeleteUser(c *gin.Context) {
	// Get Data if Exist
	db := c.MustGet("db").(*gorm.DB)
	var data models.User
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!!"})
		return
	}

	db.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
