package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qullDev/BookStore-API/internal/controllers"
)

func PaymentRoutes(r *gin.RouterGroup) {
	payments := r.Group("/payments")
	{
		payments.GET("/", controllers.GetPayments)
		payments.GET("/:id", controllers.GetPaymentByID)
		payments.POST("/", controllers.CreatePayment)
		payments.PUT("/:id", controllers.UpdatePayment)
		payments.DELETE("/:id", controllers.DeletePayment)
	}
}
