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
