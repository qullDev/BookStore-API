package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/qullDev/BookStore-API/internal/config"
	"github.com/qullDev/BookStore-API/internal/models"
	"github.com/qullDev/BookStore-API/internal/routes"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Initialize database connection
	config.InitDB()

	// Auto-migrate database schemas
	config.DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Category{})

	// Initialize Gin router
	router := gin.Default()

	// Set up API routes
	v1 := router.Group("/api/v1")
	{
		routes.AuthorRoutes(v1)
	}

	router.Run(":8080")

}
