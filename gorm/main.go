package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Name:  "Notebook",
		Price: 7899.96,
	})

	products := []Product{
		{Name: "Mouse", Price: 76.54},
		{Name: "Keyboard", Price: 200},
	}

	db.Create(&products)
}
