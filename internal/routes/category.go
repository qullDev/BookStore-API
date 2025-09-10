package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qullDev/BookStore-API/internal/controllers"
)

func CategoryRoutes(r *gin.RouterGroup) {
	category := r.Group("/categories")
	{
		category.GET("/", controllers.GetCategories)
		category.GET("/:id", controllers.GetCategoryByID)
		category.POST("/", controllers.CreateCategory)
		category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
	}
}
