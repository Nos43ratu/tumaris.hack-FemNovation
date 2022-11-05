package order

import "tumaris.hack-FemNovation/backend/internal/models"

type Service interface {
	GetAll(user *models.User) ([]*models.Order, error)
	Create(order *models.Order) error
	Update(order *models.Order) error
}
