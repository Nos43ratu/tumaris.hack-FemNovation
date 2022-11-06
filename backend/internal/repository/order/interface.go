package order

import "tumaris.hack-FemNovation/backend/internal/models"

type Order interface {
	GetByID(orderID string) (*models.Order, error)
	GetAll(userID string) ([]*models.Order, error)
	Create(order *models.Order) error
	Update(orderID string, order *models.Order) error
}
