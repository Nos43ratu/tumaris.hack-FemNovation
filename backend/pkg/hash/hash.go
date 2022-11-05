package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHasher struct for crypto hasher
type BcryptHasher struct {
	PasswordPepper string
}

// PasswordHasher interface for password works
type PasswordHasher interface {
	Hash(password string, cost int) (string, error)
	Compare(hash, password string) error
}

// NewByCryptHasher creates new hasher with papper
func NewByCryptHasher(pepper string) *BcryptHasher {
	return &BcryptHasher{
		PasswordPepper: pepper,
	}
}

// Hash get passwod and his cost and creates new hash for them
func (h *BcryptHasher) Hash(password string, cost int) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password+h.PasswordPepper), cost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// Compare gets pass and hash and compares them
func (h *BcryptHasher) Compare(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, append(password, []byte(h.PasswordPepper)...))
}
