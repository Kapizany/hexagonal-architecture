package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Kapizany/hexagonal-architecture/adapters/db"
	"github.com/Kapizany/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float64, "status" string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values ("p1", "Product test 1", 0.0, "disabled"),
				("p2", "Product test 2", 9.9, "disabled"), ("p3", "Product test 3", 88.8, "disabled");
			`
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
	product, err := productDb.Get("p1")
	require.Nil(t, err)
	require.Equal(t, "Product test 1", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product test 10"
	product.Price = 999.99

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())
}
