package order

import (
	"rest_api_order/internal/repository/database"
	"rest_api_order/internal/repository/models/item"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint        `gorm:"primaryKey"`
	CustomerName string      `json:"customer_name"`
	Items        []item.Item `json:"items"`
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var db *gorm.DB = database.New()

func GetAllData() *[]Order {
	var orders []Order
	db.Preload("Items").Find(&orders)
	return &orders
}

func GetSingleData(id uint) *Order {
	var order Order
	db.Preload("Items").Find(&order, "id = ?", id)
	return &order
}

func InsertData(newOrder *Order) {
	db.Create(newOrder)
}
