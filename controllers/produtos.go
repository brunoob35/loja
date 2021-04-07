package controllers

import (
	"github.com/brunoob35/loja/models"
	"net/http"
	"html/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func  Index(w http.ResponseWriter, r *http.Request)  {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

