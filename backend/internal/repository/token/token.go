package token

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type TokenRepo struct {
	sqlite  *sql.DB
	logger  *zap.SugaredLogger
	timeout time.Duration

	privateKey []byte
	publicKey  []byte
}

func NewTokenRepo(logger *zap.SugaredLogger, sqlite *sql.DB, sqliteTimeout time.Duration, privateKey, publicKey []byte) Token {
	return &TokenRepo{
		logger:     logger,
		sqlite:     sqlite,
		timeout:    sqliteTimeout,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (t *TokenRepo) CreateToken(ttl time.Duration, opts map[string]interface{}, role string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(t.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)

	for key, value := range opts {
		claims[key] = value
	}

	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix() // The time at which the token was issued.
	claims["role"] = role

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (t *TokenRepo) SetSession(session *models.Session) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	query := `INSERT INTO tokens (user_id, access_token, refresh_token, refresh_expire) VALUES ($1, $2, $3, $4)`
	_, err := t.sqlite.ExecContext(ctx, query, session.UserID, session.AccessToken, session.RefreshToken, session.RefreshExpire)
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenRepo) GetSession(refreshToken string) (*models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	s := &models.Session{}
	query := `SELECT * FROM tokens WHERE refresh_token = $1`
	err := t.sqlite.QueryRowContext(ctx, query, refreshToken).Scan(&s.ID, &s.UserID, &s.RefreshToken, &s.AccessToken, &s.RefreshExpire)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}

	return s, nil
}

func (t *TokenRepo) DeleteSession(refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	query := `DELETE FROM tokens WHERE access_token = $1`
	_, err := t.sqlite.ExecContext(ctx, query, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenRepo) ParseTokenWithoutClaims(token string) error {
	key, err := jwt.ParseRSAPublicKeyFromPEM(t.publicKey)
	if err != nil {
		return fmt.Errorf("parse without claims: parse key: %w", err)
	}

	_, err = (&jwt.Parser{SkipClaimsValidation: false}).Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method validate: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

// ParseTokenWithClaims parses tokens claims
func (t *TokenRepo) ParseTokenWithClaims(token string) (jwt.MapClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(t.publicKey)
	if err != nil {
		return nil, fmt.Errorf("parse with claims: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method validate: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims, nil
}

func (t *TokenRepo) VerifyToken(token string) (bool, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(t.publicKey)
	if err != nil {
		return false, err
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Println("split token length != 3")
		return false, fmt.Errorf("invalid: token length")
	}

	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], key)
	if err != nil {
		return false, fmt.Errorf("couldn't verify validate: %w", err)
	}
	return true, nil
}
