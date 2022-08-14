package main

import (
	"api_tarefas/database"
	"api_tarefas/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Servidor Online.")
	routes.HandleRequest()
}
