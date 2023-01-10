package main

import (
	"net/http"

	"web_com_golang-aula_5/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
