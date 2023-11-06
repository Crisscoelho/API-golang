package controllers

import (
	"API/database"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodasAccounts(c *gin.Context) {
	var accounts []models.Account
	database.DB.Find(&accounts)
	c.JSON(200, accounts)

}


func NewAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&account)
	c.JSON(http.StatusOK, account)
}


func BuscaAccountPorID(c *gin.Context) {
	var accountid models.Account
	id := c.Params.ByName("id")
	database.DB.First(&accountid, id)
	if accountid.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Account não encontrada!"})
		return
	}
	c.JSON(http.StatusOK, accountid)
}

func DeletaAccount(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id")
	database.DB.Delete(&account, id)
	c.JSON(http.StatusOK, gin.H{"data": "Account deletada com sucesso!"})
}

func EditaAccount(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id")
	database.DB.First(&account, id)

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&account).UpdateColumns(account)
	c.JSON(http.StatusOK, account)
}

