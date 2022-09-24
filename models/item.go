package models

import (
	"time"
)

type Item struct {
	ItemID      uint      `gorm:"primaryKey" json:"items_id"`
	ItemCode    int       `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     int       `json:"orders_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemInput struct {
	ItemCode    int    `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	OrderID     int    `json:"orders_id"`
	ItemID      int    `json:"item_id"`
}
