package auth

import "tumaris.hack-FemNovation/backend/internal/models"

type Auth interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int) (*models.User, error)
}
