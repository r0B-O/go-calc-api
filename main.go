package main

import (
	"go-calc-api/routes"
	"log"
	"net/http"
)

func main() {
	routes.SetRoutes()
	log.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
