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
	r.HandleFunc("/receitas", controllers.Receitas).Methods("Get")
	r.HandleFunc("/receitas/{id}", controllers.RetornaUmaReceita).Methods("Get")
	r.HandleFunc("/receitas", controllers.CadastraReceita).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
