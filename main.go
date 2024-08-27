package main

import (
	"database/sql"
	"log"

	db2 "github.com/gabrielgsd/go-hexagonal/adapters/db"
	"github.com/gabrielgsd/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, err := productService.Create("Product Test", 10)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	productService.Enable(product)
}
