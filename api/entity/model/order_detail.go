package model

import "github.com/jinzhu/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID		int		`json:"order_id"`
	ProductID	int		`json:"product_id"`
	Quantity	int		`json:"quantity"`
}
