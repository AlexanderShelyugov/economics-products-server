package model

type ProductType struct {
	tableName struct{} `pg:"PRODUCT_TYPES"`
	Id int
	Name string
}

type Product struct {
	tableName struct{} `pg:"PRODUCTS"`
	Id int
	Uuid string `pg:unique`
	Name string
	ProductType int
}