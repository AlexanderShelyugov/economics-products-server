package model

import (
	"github.com/go-pg/pg/v9"
	"economics/products/internal/config"
)

func GetProducts() []Product {
	config := config.GetDBConfig()
    db := pg.Connect(&pg.Options{
		User: config.User,
		Password: config.Password,
		Database: config.Database,
    })
    defer db.Close()

	var products []Product
	err := db.Model(&products).
		Join("JOIN PRODUCT_TYPES t ON t.ID = product.PRODUCT_TYPE").
		Select()
	if err != nil {
		panic(err)
	}
	return products
}