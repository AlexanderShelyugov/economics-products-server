package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"economics/products/internal/config"
)

func Init() (*gorm.DB, error) {
	config := config.GetDBConfig()
	url := fmt.Sprint(
		"host=", config.Host,
		" port =", config.Port,
		" user=", config.User,
		" dbname=", config.Database,
		" password=", config.Password,
		" sslmode=disable",
	)
	db, err := gorm.Open("postgres", url)
	db.LogMode(true)
	if  err != nil {
		return nil, err
	}
	return db, nil
}