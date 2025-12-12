package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ProductName string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`
	Quantity    int16  `gorm:"default:0"`
}

type UserData struct {
	gorm.Model

	Name  string `gorm:"type:varchar(100)" json:"name"`
	Email string `gorm:"type:varchar(100);uniqueIndex" json:"email"`

	Age int16 `gorm:"default:0"`
}
