package main

import (
	"API/database"

	"API/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
