package handlers

import (
	"encoding/json"
	"go-calc-api/calc"
	"net/http"
	"strconv"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	queryParam1 := queryParams.Get("num1")
	queryParam2 := queryParams.Get("num2")

	num1, err := strconv.Atoi(queryParam1)
	if err != nil {
		http.Error(w, "You sent an invalid first number", http.StatusBadRequest)
	}

	num2, err := strconv.Atoi(queryParam2)
	if err != nil {
		http.Error(w, "You sent an invalid second number", http.StatusBadRequest)
	}

	sum := calc.Sum(num1, num2)

	response := map[string]int{"sum": sum}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
