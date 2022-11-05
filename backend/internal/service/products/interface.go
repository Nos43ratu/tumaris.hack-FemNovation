package products

import (
	"tumaris.hack-FemNovation/backend/internal/models"
)

type Service interface {
	CreateProduct(product *models.Product) (int, error)
	UpdateProduct(product *models.Product) (int, error)
}
