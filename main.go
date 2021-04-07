package main

import (
	"net/http"
	"github.com/brunoob35/loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

