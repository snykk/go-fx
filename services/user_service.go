package services

import (
	"github.com/snykk/go-fx/models"
	"github.com/snykk/go-fx/repositories"
)

type IUserService interface {
	GetUser(id int) (*models.User, error)
}

type userServiceImpl struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) IUserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) GetUser(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
