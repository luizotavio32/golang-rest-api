package main

import (
	"fmt"

	"github.com/luizotavio32/golang-rest-api/database"
	"github.com/luizotavio32/golang-rest-api/models"
	"github.com/luizotavio32/golang-rest-api/routes"
)


func main() {

	models.Personalidades = []models.Personalidade{
		{Id:1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id:2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Inciando servidor rest")
	routes.HandleRequest()
}