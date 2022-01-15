package main

import (
	"html/template"
	"mvc/models"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)

}
