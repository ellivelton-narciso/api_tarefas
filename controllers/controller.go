package controllers

import (
	"api_tarefas/database"
	"api_tarefas/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func Receitas(w http.ResponseWriter, r *http.Request) {
	var p []models.Receitas
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaReceita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var receita models.Receitas
	database.DB.First(&receita, id)
	json.NewEncoder(w).Encode(receita)

}

func Despesas(w http.ResponseWriter, r *http.Request) {
	var d []models.Despesas
	database.DB.Find(&d)
	json.NewEncoder(w).Encode(d)
}

func RetornaUmaDespesa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var despesa models.Receitas
	database.DB.First(&despesa, id)
	json.NewEncoder(w).Encode(despesa)
}
