package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qullDev/BookStore-API/internal/config"
	"github.com/qullDev/BookStore-API/internal/models"
)

func GetAuthors(c *gin.Context) {
	// Implementation for getting all authors
	var authors []models.Author

	if err := config.DB.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func GetAuthorByID(c *gin.Context) {
	// Implementation for getting an author by ID
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateAuthor(c *gin.Context) {
	// Implementation for creating a new author
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{Name: input.Name}
	if err := config.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

type UpdateAuthorInput struct {
	Name string `json:"name"`
}

func UpdateAuthor(c *gin.Context) {
	// Implementation for updating an existing author
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	var input UpdateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		author.Name = input.Name
	}

	if err := config.DB.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
	// Implementation for deleting an author
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	if err := config.DB.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
