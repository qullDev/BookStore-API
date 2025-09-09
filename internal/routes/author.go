package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qullDev/BookStore-API/internal/controllers"
)

func AuthorRoutes(api *gin.RouterGroup) {

	api.GET("/authors", controllers.GetAuthors)
	api.GET("/authors/:id", controllers.GetAuthorByID)
	api.POST("/authors", controllers.CreateAuthor)
	api.PUT("/authors/:id", controllers.UpdateAuthor)
	api.DELETE("/authors/:id", controllers.DeleteAuthor)
}
