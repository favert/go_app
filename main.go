package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/favert/go_app/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
