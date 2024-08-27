package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/gabrielgsd/go-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createTable := `
		CREATE TABLE products (
			"id" string PRIMARY KEY,
			"name" string,
			"price" FLOAT,
			"status" string
		);
	`
	stmt, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product 1", 10, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
