package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	BuyerID 	int		`json:"buyer_id"`
	Quantity 	int		`json:"quantity"`
	StoreID		int		`json:"store_id"`
}
