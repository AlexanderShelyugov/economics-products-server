package main

import (
	"encoding/json"
	"net/http"
	"economics/products/internal/model"
)

func main() {
	model.GetProducts()
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8100", nil)
}


func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := model.GetProducts()
	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}