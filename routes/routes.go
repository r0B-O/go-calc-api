package routes

import (
	"go-calc-api/handlers"
	"net/http"
)

func SetRoutes() {
	http.HandleFunc("/sum", handlers.SumHandler)
}
