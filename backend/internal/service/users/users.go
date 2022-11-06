package users

import (
	"log"

	"tumaris.hack-FemNovation/backend/internal/models"
	"tumaris.hack-FemNovation/backend/internal/repository/users"
)

type UserService struct {
	Users users.Users
}

func NewUserService(users users.Users) Service {
	return &UserService{
		Users: users,
	}
}

func (u *UserService) GetByEmail(email string) (*models.UserInfo, error) {
	log.Printf("[%s]", email)
	return u.Users.GetByEmail(email)
}
