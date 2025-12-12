package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ProductName string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`
	Quantity    int16  `gorm:"default:0"`
}
