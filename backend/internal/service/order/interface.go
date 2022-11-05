package order

import "tumaris.hack-FemNovation/backend/internal/models"

type Service interface {
	GetByID(orderID string) (*models.Order, error)
	GetAll(user *models.User) ([]*models.Order, error)
	Create(order *models.Order) error
	Update(orderID string, order *models.Order) error
}
