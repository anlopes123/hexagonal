package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db1 "github.com/anlopes123/hexagonal/adapter/db"
	"github.com/anlopes123/hexagonal/application"
)

func main(){

	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db1.NewProductDb(db)

	productService:= application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Exemple", 30)
	productService.Enable(product)


	
}