package main

import (
	"net/http"

	routes "github.com/gophercises/prob_3"
)

func main() {
	routes.HandleRoutes()
	http.ListenAndServe(":8080", nil)
}
