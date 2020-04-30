package model

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"economics/products/internal/config"
)

func runMigrations(config *config.DBConfig) {
	url := fmt.Sprint("postgres://", config.User, ":", config.Password, "@", config.Host, ":", config.Port, "/", config.Database, "?sslmode=disable")
	m, err := migrate.New("file://internal/model/migrations", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func Init() (*gorm.DB, error) {
	config := config.GetDBConfig()
	runMigrations(&config)
	url := fmt.Sprint(
		"host=", config.Host,
		" port =", config.Port,
		" user=", config.User,
		" dbname=", config.Database,
		" password=", config.Password,
		" sslmode=disable",
	)
	db, err := gorm.Open("postgres", url)
	if  err != nil {
		return nil, err
	}
	return db, nil
}