package order

import "tumaris.hack-FemNovation/backend/internal/models"

type Order interface {
	GetAll() ([]*models.Order, error)
	Create(order *models.Order) error
	Update(order *models.Order) error
}
