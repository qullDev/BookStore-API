package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qullDev/BookStore-API/internal/config"
	"github.com/qullDev/BookStore-API/internal/models"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := config.DB.Preload("Author").Preload("Category").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByQuery(c *gin.Context) {
	title := c.Query("title")
	authorName := c.Query("author")
	categoryName := c.Query("category")

	var books []models.Book
	query := config.DB.Preload("Author").Preload("Category")

	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if authorName != "" {
		query = query.Joins("JOIN authors ON authors.id = books.AuthorID").Where("authors.name ILIKE ?", "%"+authorName+"%")
	}
	if categoryName != "" {
		query = query.Joins("JOIN categories ON category.id = books.CategoryID").Where("Category.name ILIKE ?", "%"+categoryName+"%")
	}

	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.Preload("Author").Preload("Category").First(&book, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

type CreateBookInput struct {
	Title      string    `json:"title" binding:"required"`
	AuthorID   uuid.UUID `json:"author_id" binding:"required,uuid"`
	CategoryID uuid.UUID `json:"category_id" binding:"required,uuid"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, AuthorID: input.AuthorID, CategoryID: input.CategoryID}
	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

type UpdateBookInput struct {
	Title      string    `json:"title"`
	AuthorID   uuid.UUID `json:"author_id" binding:"omitempty,uuid"`
	CategoryID uuid.UUID `json:"category_id" binding:"omitempty,uuid"`
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook := models.Book{
		Title:      input.Title,
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
	}

	if err := config.DB.Model(&book).Updates(updatedBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
