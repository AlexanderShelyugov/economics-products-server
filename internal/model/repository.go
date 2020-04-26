package model

import (
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r ProductRepository) GetAll() []Product {
	var products []Product
	r.db.Preload("ProductType").Find(&products)
	return products
}