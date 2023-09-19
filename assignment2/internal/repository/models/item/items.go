package item

import (
	"rest_api_order/internal/repository/database"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"unique" json:"code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var db *gorm.DB = database.New()

func DoesDuplicateExist(item Item) bool {
	var tempItem Item
	db.Find(&tempItem, "code = ?", item.Code)
	if tempItem.ID == 0 {
		return false
	}
	return true
}

func DoDuplicatesExist(items []Item) bool {
	var result bool = false
	for _, item := range items {
		result = result || DoesDuplicateExist(item)
	}
	return result
}
