package server

import (
	"fmt"
	"encoding/json"
	"net/http"
	"economics/products/internal/config"
	"economics/products/internal/model"
)

type Server struct {
	productsRepo *model.ProductRepository
}

type Result struct {

}

func NewServer(r *model.ProductRepository) *Server {
	return &Server{r}
}

func (s Server) Run() {
	http.HandleFunc("/products", s.productsHandler)
	port := fmt.Sprint(":", config.GetPort())
	http.ListenAndServe(port, nil)
}

func (s Server) productsHandler(w http.ResponseWriter, r *http.Request) {
	products := s.productsRepo.GetAll()
	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}