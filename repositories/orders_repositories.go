package repositories

import (
	"golang-jwt/models"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	GetOrders() ([]models.Order, error)
	GetOrderByOrderID(orderID int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(orderID int, order models.Order) (models.Order, error)
	DeleteOrder(order models.Order) error
}

type ordersRepository struct {
	db *gorm.DB
}

func (r *ordersRepository) GetOrders() ([]models.Order, error) {
	var orders []models.Order

	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *ordersRepository) GetOrderByOrderID(orderID int) (models.Order, error) {
	var orders models.Order

	err := r.db.Preload("Items").Where("order_id = ?", orderID).First(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *ordersRepository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *ordersRepository) UpdateOrder(orderID int, order models.Order) (models.Order, error) {
	err := r.db.Preload("Items").Where("order_id = ?", orderID).Updates(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *ordersRepository) DeleteOrder(order models.Order) error {
	err := r.db.Delete(order).Error
	if err != nil {
		return err
	}

	return err
}

func ProvideOrderRepository(db *gorm.DB) *ordersRepository {
	return &ordersRepository{db}
}