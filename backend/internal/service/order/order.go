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

func NewOrderService(order order.Order, products products.Products) Service {
	return &OrderService{
		Order:    order,
		Products: products,
	}
}

func (o *OrderService) Create(order *models.Order) error {
	return o.Order.Create(order)
}

func (o *OrderService) Update(orderID string, order *models.Order) error {
	return o.Order.Update(orderID, order)
}

func (o *OrderService) GetAll(userID string) ([]*models.Order, error) {
	orders, err := o.Order.GetAll(userID)
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
