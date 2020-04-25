package model

import (
	"fmt"
	"github.com/go-pg/pg/v9"
)

func GetProducts() []Product {
    db := pg.Connect(&pg.Options{
		User: "postgres",
		Password: "postgres",
		Database: "economics-products",
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