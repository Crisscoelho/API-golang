package routes

import (
	"API/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/banks", controllers.ExibeTodosBanks) //busca todos os bancos
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/bank/:id", controllers.BuscaBankPorID)                     //busca por id
	r.GET("/account/:id", controllers.BuscaAccountPorID)               //busca por id
	r.GET("/accounts", controllers.ExibeTodasAccounts)                 // busca todas accounts
	r.GET("/customer/:id", controllers.BuscaCustomerPorID)             //busca cliente por id
	r.GET("/customers", controllers.ExibeTodosCustomers)               // busca todos clientes
	r.GET("/customer/CPF_CNPJ:CPF_CNPJ", controllers.BuscaPorCPF_CNPJ) // busca por CPF_CNPJ
	r.POST("/bank", controllers.NewBank)                               // novo banco
	r.POST("/account", controllers.NewAccount)                         // nova conta
	r.POST("/customer", controllers.NewCustomer)                       // novo cliente
	r.POST("/transaction", controllers.NewTransaction)                 // nova transacao
	r.DELETE("/account/:id", controllers.DeletaAccount)                //exclui conta
	r.DELETE("/bank/:id", controllers.DeletaBank)                      //exclui bank
	r.DELETE("/customer/:id", controllers.DeletaCustomer)              //exclui customer
	r.DELETE("/transaction/:id", controllers.DeletaTransaction)        //exclui transaction
	r.PATCH("/account/:id", controllers.EditaAccount)                  // edita account
	r.PATCH("/bank/:id", controllers.EditaBank)                        // edita bank
	r.PATCH("/customer/:id", controllers.EditaCustomer)                // edita customer
	r.PATCH("/transaction/:id", controllers.EditaTransaction)          // edita transaction
	r.Run()
}
