package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qullDev/BookStore-API/internal/controllers"
)

func AuthorRoutes(r *gin.RouterGroup) {

	r.GET("/authors", controllers.GetAuthors)
	r.GET("/authors/:id", controllers.GetAuthorByID)
	r.POST("/authors", controllers.CreateAuthor)
	r.PUT("/authors/:id", controllers.UpdateAuthor)
	r.DELETE("/authors/:id", controllers.DeleteAuthor)
}
