package services

import (
	"errors"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	//Validation logic in the future
	if user.Username == nil || *user.Username == "" {
		return errors.New("username cannot be empty")
	}
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.repo.FindByID(id)
}

func (s *userService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.Delete(id)
}
