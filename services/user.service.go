package services

import (
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s UserService) Create(model *models.User) (*models.User, error) {
	return s.repository.Create(model)
}

func (s UserService) GetByEmail(email string) (*models.User, error) {
	return s.repository.GetByEmail(email)
}

func (s UserService) GetByUuid(uuid string) (*models.User, error) {
	return s.repository.GetByUuid(uuid)
}