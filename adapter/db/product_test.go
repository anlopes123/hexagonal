package db_test

import  (
	"testing"
	"database/sql"
	"github.com/anlopes123/hexagonal/adapter/db"
	"github.com/stretchr/testify/require"	
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table:= `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`
	stmt, err := db.Prepare(table)
	if (err != nil) {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	product:= `INSERT INTO products VALUES("abc", "Product test", 0, "disable"0)`
	stmt, err := db.Prepare(product)
	if (err != nil) {
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
	require.Equal(t, "Product test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disable", product.GetStatus())



}