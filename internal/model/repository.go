package repository

import (
	"os"
	"github.com/jackc/pgx/v4"
)

const (
	querySelectAll = "SELECT * FROM PRODUCTS"
)

func GetProducts() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), querySelectAll)
	for rows.Next() {
		var id uint
		var name string
		var productType uint
		var weight float

		err := rows.Scan(&id, &name, &productType, &weight)
		
	}

	fmt.Println(name, weight)
}