package model

type ProductType struct {
	id uint,
	name string
}

type Product struct {
	id uint,
	name string,
	productType ProductType,
	weight float
}