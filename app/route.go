package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jojiejo/hibank-loan/app/handlers"
)

func GenerateRoute(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/loan", handlers.CreateLoan)
}
