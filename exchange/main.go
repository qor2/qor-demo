package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jinzhu/gorm"
	"github.com/qor/exchange"
	"github.com/qor/exchange/backends/csv"
	"github.com/qor/qor"
)

// Product for exchange demo
type Product struct {
	gorm.Model

	Code  string
	Name  string
	Price float64
}



func main(){
	DB, _ := gorm.Open("sqlite3", "demo.db")
	defer DB.Close()

	DB.AutoMigrate(&Product{})

	// before export, we should add a new product
	DB.Create(&Product{Name:"product", Code: "mycode", Price: 42.0})

	product := exchange.NewResource(&Product{}, exchange.Config{PrimaryField: "Code"})
	product.Meta(&exchange.Meta{Name: "Code"})
	product.Meta(&exchange.Meta{Name: "Name"})
	product.Meta(&exchange.Meta{Name: "Price"})

	ctx := &qor.Context{DB: DB}

	// callbacks could omit
	product.Export(csv.New("products.csv"), ctx, func(progress exchange.Progress) error {
		fmt.Printf("%v/%v Exporting product %v\n", progress.Current, progress.Total, progress.Value.(*Product).Code)
		return nil
	})
}
