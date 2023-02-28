package db_test

import  (
	"testing"
	"database/sql"
	/*"github.com/mattn/go-sqlite3"*/
	_ "github.com/mattn/go-sqlite3"
	"github.com/anlopes123/hexagonal/adapter/db"
	"github.com/stretchr/testify/require"	
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	if Db == nil {
		log.Fatal("Db está nil")
	}
	createTable(Db)
	createProduct(Db)
}

func createTable(db1 *sql.DB) {
	table:= `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`
	stmt, err := db1.Prepare(table)
	if (err != nil) {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db1 *sql.DB) {
	product:= `INSERT INTO products VALUES("abc", "Product test", 0, "disabled")`
	stmt, err := db1.Prepare(product)
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
	require.Equal(t, "disabled", product.GetStatus())



}