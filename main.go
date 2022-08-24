package main

import (
	"fmt"

	"github.com/luizotavio32/golang-rest-api/models"
	"github.com/luizotavio32/golang-rest-api/routes"
)


func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	
	fmt.Println("Inciando servidor rest")
	routes.HandleRequest()
}