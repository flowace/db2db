package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product -- Represents a product
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// TableName setting the table name
func (Product) TableName() string {
	return "allProducts"
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Product{})
	// db.CreateTable(&Mayukh{})
	db.Create(&Product{Code: "L1212", Price: 1000})
	// db.Commit()
	// Reading the database
	var product Product
	db.First(&product, 1)
	fmt.Println(product.Code)
	fmt.Println(product.Price)
}
