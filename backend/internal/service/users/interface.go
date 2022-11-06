package users

import "tumaris.hack-FemNovation/backend/internal/models"

type Service interface {
	GetByEmail(email string) (*models.UserInfo, error)
}
