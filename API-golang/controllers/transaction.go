package controllers

import (
	"API/database"
	"API/models"
	"fmt"
	"net/http"

	_ "gorm.io/driver/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExibeTodasTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(200, transactions)

}

func NewTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&transaction)
	c.JSON(http.StatusOK, transaction)
}

func BuscaTransactionPorID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)
	if transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "transaction não encontrada!"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func DeletaTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.Delete(&transaction, id)
	c.JSON(http.StatusOK, gin.H{"data": "Transaction deletado com sucesso!"})
}

func EditaTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&transaction).UpdateColumns(transaction)
	c.JSON(http.StatusOK, transaction)
}

//corrigir regra de negocio

func Transaction(DB *gorm.DB, customer *models.Customer) (n models.Transaction) {
	fmt.Println(customer, "Seja benvindo ao Bank")
	var account1 models.Account
	var trans1 models.Transaction
	var customer1 models.Customer
	trans1.FkAccountId = customer1.CustomerId

	switch customer1.Check {
	case 1:
		trans1.RunningBalance = account1.CurrentBalance - customer1.FkCurrentBalance
		trans1.DebitedAmount = customer1.FkCurrentBalance
	case 2:
		trans1.RunningBalance = account1.CurrentBalance + customer1.FkCurrentBalance
		trans1.CreditedAmount = customer1.FkCurrentBalance
	}
	if trans1.RunningBalance < 0 {
		fmt.Println("Balance too low for transaction")
		return trans1
	}
	trans1.OtherPartyAccountNumber = customer1.FkCurrentBalance
	fmt.Println(trans1)

	return
}
