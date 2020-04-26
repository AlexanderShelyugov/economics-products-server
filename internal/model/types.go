package model

type ProductType struct {
	Id uint `gorm:"primary_key" json:"-"`
	Uuid string `gorm:"not null;unique" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (ProductType) TableName() string {
	return "product_types"
}

type Product struct {
	Id uint `gorm:"primary_key" json:"-"`
	Uuid string `gorm:"not null;unique" json:"id"`
	Name string `gorm:"not null" json:"name"`
	ProductTypeId uint `json:"-"`
	ProductType ProductType `json:"type"`
	Weight float64 `json:"weight"`
}

func (Product) TableName() string {
	return "products"
}