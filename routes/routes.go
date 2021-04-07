package routes

import (
	"github.com/brunoob35/loja/controllers"
	"net/http"
)

func CarregaRotas()  {
	http.HandleFunc("/", controllers.Index)
}

