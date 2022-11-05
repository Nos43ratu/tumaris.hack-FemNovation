package order

import "tumaris.hack-FemNovation/backend/internal/models"

type Service interface {
	Create(order *models.Order) error
}
