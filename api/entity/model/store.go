package model

import "github.com/jinzhu/gorm"

type Store struct {
	gorm.Model
	StoreName 	string	`json:"store_name"`
	UserID 		int	`json:"user_id"`
}
