package models

import (
	"time"
)

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"orders_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OrderInput struct {
	CustomerName string      `json:"customer_name" binding:"required"`
	OrderedAt    time.Time   `json:"ordered_at" binding:"required"`
	Items        []ItemInput `json:"items" binding:"required"`
}
