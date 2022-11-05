package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type Token interface {
	CreateToken(ttl time.Duration, opts map[string]interface{}, role string) (string, error)
	GetSession(refreshToken string) (*models.Session, error)
	SetSession(session *models.Session) error
	DeleteSession(refreshToken string) error
	VerifyToken(token string) (bool, error)
	ParseTokenWithClaims(token string) (jwt.MapClaims, error)
	ParseTokenWithoutClaims(token string) error
}
