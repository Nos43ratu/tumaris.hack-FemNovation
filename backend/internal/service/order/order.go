package order

import (
	"tumaris.hack-FemNovation/backend/internal/models"
	"tumaris.hack-FemNovation/backend/internal/repository/order"
	"tumaris.hack-FemNovation/backend/internal/repository/products"
)

type OrderService struct {
	Order    order.Order
	Products products.Products
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
	orders, err := o.Order.GetAll(user)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		order.Products, err = o.Products.GetProductByID(order.ProductID)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (o *OrderService) GetByID(orderID string) (*models.Order, error) {
	order, err := o.Order.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	order.Products, err = o.Products.GetProductByID(order.ProductID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
