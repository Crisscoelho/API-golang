package controllers

import (
	"API/database"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosBanks(c *gin.Context) {
	var banks []models.Bank
	database.DB.Find(&banks)
	c.JSON(200, banks)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": nome + ": " + "Seja benvindo ao nosso Bank!",
	})
}

func NewBank(c *gin.Context) {
	var bank models.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&bank)
	c.JSON(http.StatusOK, bank)
}

func BuscaBankPorID(c *gin.Context) {
	var bank models.Bank
	id := c.Params.ByName("id")
	database.DB.First(&bank, id)
	if bank.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Bank não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, bank)
}

func DeletaBank(c *gin.Context) {
	var bank models.Bank
	id := c.Params.ByName("id")
	database.DB.Delete(&bank, id)
	c.JSON(http.StatusOK, gin.H{"data": "Bank deletado com sucesso!"})
}

func EditaBank(c *gin.Context) {
	var bank models.Bank
	id := c.Params.ByName("id")
	database.DB.First(&bank, id)

	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&bank).UpdateColumns(bank)
	c.JSON(http.StatusOK, bank)
}
