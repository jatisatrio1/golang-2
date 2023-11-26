package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Item_code   string `json:"itemCode" form:"itemCode" query:"itemCode" gorm:"not null;varchar(200)"`
	Description string `json:"description" form:"description" query:"description" gorm:"not null;varchar(200)"`
	Quantity    int    `json:"quantity" form:"quantity" query:"quantity" gorm:"not null"`
	Order_id    int    `gorm:"not null"`
}

type Order struct {
	gorm.Model
	Customer_name string `json:"customer_name" form:"customer_name" query:"customer_name" gorm:"not null;varchar(200)"`
	Items         []Item
	Ordered_at    time.Time `gorm:"not null"`
}

type Response struct {
	Message string
	Success bool
	Data    interface{}
}
