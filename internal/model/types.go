package model

type ProductType struct {
	Id int
	Uuid string `gorm:"not null;unique"`
	Name string `gorm:"not null"`
}

func (ProductType) TableName() string {
	return "product_types"
}

type Product struct {
	Id int
	Uuid string `gorm:"not null;unique"`
	Name string `gorm:"not null"`
	Type ProductType `gorm:"foreignkey:PRODUCT_TYPE"`
}

func (Product) TableName() string {
	return "products"
}