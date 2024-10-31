package routes

import (
	"net/http"

	"github.com/favert/go_app/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
