package services

import (
	"golang-jwt/models"
	"golang-jwt/repositories"
)

type OrdersService interface {
	GetOrders() ([]models.Order, error)
	GetOrder(orderID int) (models.Order, error)
	CreateOrder(input models.OrderInput) (models.Order, error)
	UpdateOrder(orderID int, input models.OrderInput) (models.Order, error)
	DeleteOrder(orderID int) error
}

type orderService struct {
	orderRepository repositories.OrdersRepository
	itemsRepository repositories.ItemRepository
}

func (s *orderService) GetOrders() ([]models.Order, error) {
	orders, err := s.orderRepository.GetOrders()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *orderService) GetOrder(orderID int) (models.Order, error) {
	order, err := s.orderRepository.GetOrderByOrderID(orderID)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *orderService) CreateOrder(input models.OrderInput) (models.Order, error) {
	var order models.Order
	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt

	order, err := s.orderRepository.CreateOrder(order)
	if err != nil {
		return order, err
	}

	for _, item := range input.Items {
		newItem, err := s.itemsRepository.CreateItem(models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     int(order.OrderID),
		})
		if err != nil {
			return order, err
		}
		order.Items = append(order.Items, newItem)
	}

	return order, nil
}

func (s *orderService) UpdateOrder(orderID int, input models.OrderInput) (models.Order, error) {
	order, err := s.orderRepository.GetOrderByOrderID(orderID)
	if err != nil {
		return order, err
	}

	for _, itemInput := range input.Items {
		item, err := s.itemsRepository.GetItemByItemID(itemInput.ItemID)
		if err != nil {
			return order, err
		}

		item.ItemCode = itemInput.ItemCode
		item.Description = itemInput.Description
		item.Quantity = itemInput.Quantity
		_, err = s.itemsRepository.UpdateItem(itemInput.ItemID, item)
		if err != nil {
			return order, err
		}
	}

	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt

	order, err = s.orderRepository.UpdateOrder(orderID, order)
	if err != nil {
		return order, err
	}

	// Get order
	order, err = s.orderRepository.GetOrderByOrderID(orderID)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *orderService) DeleteOrder(orderID int) error {
	order, err := s.orderRepository.GetOrderByOrderID(orderID)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err := s.itemsRepository.DeleteItem(item)
		if err != nil {
			return err
		}
	}

	err = s.orderRepository.DeleteOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func ProvideOrderService(orderRepository repositories.OrdersRepository, itemsRepository repositories.ItemRepository) *orderService {
	return &orderService{orderRepository, itemsRepository}
}