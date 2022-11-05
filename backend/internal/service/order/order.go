package order

import (
	"tumaris.hack-FemNovation/backend/internal/models"
	"tumaris.hack-FemNovation/backend/internal/repository/order"
)

type OrderService struct {
	Order order.Order
}

func NewOrderService(order order.Order) Service {
	return &OrderService{
		Order: order,
	}
}

func (o *OrderService) Create(order *models.Order) error {
	return o.Order.Create(order)
}

func (o *OrderService) Update(orderID string, order *models.Order) error {
	return o.Order.Update(orderID, order)
}

func (o *OrderService) GetAll(user *models.User) ([]*models.Order, error) {
	return o.Order.GetAll(user)
}

func (o *OrderService) GetByID(orderID string) (*models.Order, error) {
	return o.Order.GetByID(orderID)
}
