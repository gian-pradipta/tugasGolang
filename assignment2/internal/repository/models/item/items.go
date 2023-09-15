package item

import (
	"rest_api_order/internal/repository/database"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var db *gorm.DB = database.New()
