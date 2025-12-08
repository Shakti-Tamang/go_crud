package models

type Item struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`

	ProductName string `json:"product_name"`

	Description string `json:"description"`

	Quantity int16 `json:"quantity"`
}
