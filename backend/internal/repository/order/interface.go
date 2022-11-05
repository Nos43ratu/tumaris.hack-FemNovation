package order

import "tumaris.hack-FemNovation/backend/internal/models"

type Order interface {
	Create(order *models.Order) error
}
