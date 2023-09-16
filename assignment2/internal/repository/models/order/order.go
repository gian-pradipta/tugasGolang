package order

import (
	"errors"
	"fmt"
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

func DeleteData(id uint) error {
	var order Order
	var err error
	db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		err = errors.New("Data Not Found")
		return err
	}
	db.Delete(&Order{}, id)
	return err
}

func UpdateAnEntireOrder(id uint, newOrder *Order) error {
	var err error
	var order Order
	fmt.Println(newOrder)
	db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		err = errors.New("Data not found")
		return err
	}
	db.Model(&order).Updates(newOrder)
	order.Items = newOrder.Items
	db.Save(&order)
	return err

}

func UpdatePartOfOrder(id uint, order *Order) {

}
