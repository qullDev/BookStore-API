package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qullDev/BookStore-API/internal/controllers"
)

func BookRoutes(r *gin.RouterGroup) {
	books := r.Group("/books")
	{
		books.GET("/", controllers.GetBooks)
		books.GET("/:id", controllers.GetBookByID)
		books.POST("/", controllers.CreateBook)
		books.PUT("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}