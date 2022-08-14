package models

type Receitas struct {
	Id        int     `json:"id"`
	Descricao string  `json:"descricao"`
	Valor     float32 `json:"valor"`
	Data      string  `json:"data"`
}

var ReceitasTudo []Receitas
