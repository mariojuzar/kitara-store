package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName		string	`json:"product_name"`
	CurrentStock 	int		`json:"current_stock"`
	StoreID 		uint	`json:"store_id"`
}
