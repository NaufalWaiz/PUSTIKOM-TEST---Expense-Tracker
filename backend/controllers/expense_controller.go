package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"expense-tracker/models"
	"expense-tracker/services"
)

func GetExpenses(c *gin.Context) {
	category := c.Query("category")
	expenses, err := services.GetAllExpenses(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateExpense(expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Expense created"})
}

func UpdateExpense(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateExpense(id, expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense updated"})
}

func DeleteExpense(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	err := services.DeleteExpense(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted"})
}
