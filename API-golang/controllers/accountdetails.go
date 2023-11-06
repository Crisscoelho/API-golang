package controllers

import (
	"API/database"
	"API/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ExibeAccountDetails(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id")
	database.DB.Where(&account, id)
	fmt.Println(account)

}
