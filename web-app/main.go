package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var tmplt = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, kkkkk", Preco: 32, Quantidade: 4},
		{Nome: "Calça", Descricao: "Branca", Preco: 98, Quantidade: 12},
		{Nome: "Camisa Social", Descricao: "Preta texturizada", Preco: 140.99, Quantidade: 2},
		{Nome: "Bermuda", Descricao: "Jeans claro", Preco: 56.50, Quantidade: 7},
	}

	tmplt.ExecuteTemplate(w, "Index", produtos)
}
