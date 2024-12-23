package repositories

import (
	"github.com/snykk/go-fx/models"
)

type IUserRepository interface {
	GetUserByID(id int) (*models.User, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() IUserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) GetUserByID(id int) (*models.User, error) {
	return &models.User{ID: id, Name: "John Doe"}, nil
}
