package main

import (
	"economics/products/internal/config"
	"economics/products/internal/model"
	"economics/products/internal/server"
)

func main() {
	config.Init()
	db, err := model.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productsRepository := model.NewProductRepository(db)
	server := server.NewServer(productsRepository)
	server.Run()
}