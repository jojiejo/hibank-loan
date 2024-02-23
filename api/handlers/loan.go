package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jojiejo/hibank-loan/api/services"
)

func (server *Server) CreateLoan(c *gin.Context) {
	var loan services.Loan
	if err := c.BindJSON(&loan); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	installments := loan.CalculateInstallments()

	c.JSON(http.StatusOK, installments)
}
