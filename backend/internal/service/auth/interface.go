package auth

import (
	"github.com/dgrijalva/jwt-go"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type Service interface {
	SignIn(user *models.User) (*models.Tokens, error)
	SignOut(refreshToken string) error
	Refresh(refreshToken string) (*models.Tokens, error)
	GetUserByEmail(email string) (*models.User, error)
	ParseTokenWithClaims(token string) (jwt.MapClaims, error)
	VerifyToken(token string) (bool, error)
}
