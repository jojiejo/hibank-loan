package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jojiejo/hibank-loan/api/services"
)

func (server *Server) CreateLoan(c *gin.Context) {
	var loan services.Loan
	if err := c.BindJSON(&loan); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := server.Validate.Struct(loan)
	if err != nil {
		var errList []string
		for _, err := range err.(validator.ValidationErrors) {
			errList = append(errList, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errList,
		})
		return
	}

	installments := loan.CalculateInstallments()

	c.JSON(http.StatusOK, installments)
}
