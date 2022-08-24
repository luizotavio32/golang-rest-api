package main

import (
	"fmt"
	"github.com/luizotavio32/golang-rest-api/routes"
)


func main() {
	fmt.Println("Inciando servidor rest")
	routes.HandleRequest()
}