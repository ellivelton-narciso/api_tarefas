package routes

import (
	"api_tarefas/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	//Receita
	r.HandleFunc("/receitas", controllers.Receitas).Methods("Get")
	r.HandleFunc("/receitas/{id}", controllers.RetornaUmaReceita).Methods("Get")

	//Despesa
	r.HandleFunc("/despesas", controllers.Despesas).Methods("Get")
	r.HandleFunc("/despesas/{id}", controllers.RetornaUmaDespesa).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
