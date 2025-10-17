package routes

import (
	"expense-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/expenses", controllers.GetExpenses)
		api.POST("/expenses", controllers.CreateExpense)
		api.PUT("/expenses/:id", controllers.UpdateExpense)
		api.DELETE("/expenses/:id", controllers.DeleteExpense)
	}
}
