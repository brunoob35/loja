package routes

import (
	"github.com/brunoob35/loja/controllers"
	"net/http"
)

func CarregaRotas()  {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit )
}

