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
	OrderedAt    string      `json:"ordered_at"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var db *gorm.DB = database.New()

func isOrderExist(id uint) (*Order, error) {
	var order Order
	var err error
	db.Preload("Items").Find(&order, "id = ?", id)
	if order.ID == 0 {
		err = errors.New("Data Not Found")
		return &order, err
	}
	return &order, err
}

func GetAllData() *[]Order {
	var orders []Order
	db.Preload("Items").Find(&orders)
	return &orders
}

func GetSingleData(id uint) (*Order, error) {
	var order *Order
	var err error
	order, err = isOrderExist(id)
	if err != nil {
		return order, err
	}
	db.Preload("Items").Find(&order, "id = ?", id)
	return order, err
}

func InsertData(newOrder *Order) uint {
	db.Create(newOrder)
	return newOrder.ID
}

func DeleteData(id uint) (*Order, error) {
	var order *Order
	var err error
	order, err = isOrderExist(id)

	var deletedOrder Order = *order
	if err != nil {
		return order, err
	}
	db.Delete(order, id)
	return &deletedOrder, err
}

func UpdateOrder(id uint, newOrder *Order) (*Order, error) {
	var err error
	var order *Order
	order, err = isOrderExist(id)
	if err != nil {
		return order, err
	}

	order.CustomerName = newOrder.CustomerName
	order.OrderedAt = newOrder.OrderedAt
	for _, singleItem := range newOrder.Items {
		if err = item.UpdateItemOnCode(&singleItem); err != nil {
			fmt.Println(item.InsertData(order.ID, &singleItem))
		}
	}

	db.Save(order)
	newOrder, err = isOrderExist(order.ID)
	return newOrder, err
}
