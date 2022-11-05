package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"tumaris.hack-FemNovation/backend/internal/models"
)

type AuthService struct {
	auth  repository.Auth
	token repository.Token

	hasher     *hash.BcryptHasher
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func newAuthService(auth repository.Auth, hasher *hash.BcryptHasher, token repository.Token, accessTTL time.Duration, refreshTTL time.Duration) Service {
	return &AuthService{
		auth:       auth,
		hasher:     hasher,
		token:      token,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (a *AuthService) SignIn(input *models.User) (*models.Tokens, error) {
	user, err := a.auth.GetUserByEmail(input.Email)
	if err != nil {
		fmt.Println("!!!!!!!!!", err)
		return nil, err
	}

	fmt.Printf("%+v\n", user)
	fmt.Println(">", input.Password)
	err = a.hasher.Compare([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Println("?????????")
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, models.ErrUnauthorized
		}
		return nil, models.ErrInternalServer
	}

	claims := make(map[string]interface{})

	tokens := &models.Tokens{}
	tokens.Access, err = a.token.CreateToken(a.accessTTL, claims, user.Role)
	if err != nil {
		return nil, err
	}

	tokens.Refresh = uuid.New().String()
	refreshTTL := time.Now().Add(a.refreshTTL).Unix()
	err = a.token.SetSession(&models.Session{
		UserID:        user.ID,
		AccessToken:   tokens.Access,
		RefreshToken:  tokens.Refresh,
		RefreshExpire: refreshTTL,
	})

	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) Refresh(refreshToken string) (*models.Tokens, error) {
	session, err := a.token.GetSession(refreshToken)
	if err != nil {
		return nil, err
	}

	if session.RefreshExpire <= time.Now().Unix() {
		return nil, errors.New("refresh token is expired")
	}

	if err := a.token.DeleteSession(session.RefreshToken); err != nil {
		return nil, err
	}

	user, err := a.auth.GetUserByID(session.UserID)
	if err != nil {
		return nil, err
	}

	claims := make(map[string]interface{})

	tokens := &models.Tokens{}
	tokens.Access, err = a.token.CreateToken(a.accessTTL, claims, user.Role)
	if err != nil {
		return nil, err
	}

	tokens.Refresh = uuid.New().String()
	refreshTTL := time.Now().Add(a.refreshTTL).Unix()
	err = a.token.SetSession(&models.Session{
		UserID:        user.ID,
		AccessToken:   tokens.Access,
		RefreshToken:  tokens.Refresh,
		RefreshExpire: refreshTTL,
	})

	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) SignOut(accessToken string) error {
	if err := a.token.DeleteSession(accessToken); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return a.auth.GetUserByEmail(email)
}

func (a *AuthService) ParseTokenWithClaims(token string) (jwt.MapClaims, error) {
	return a.token.ParseTokenWithClaims(token)
}

func (a *AuthService) VerifyToken(token string) (bool, error) {
	return a.token.VerifyToken(token)
}
