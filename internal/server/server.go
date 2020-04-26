package server

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/products", s.productsHandler)
	handler := cors.Default().Handler(mux)
	port := fmt.Sprint(":", config.GetPort())
	http.ListenAndServe(port, handler)
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