package services

import (
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/utils"
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

func (s UserService) ChangePassword(model *models.User, req *models.ChangePasswordRequest) (*models.User, error) {
	// validate 
	var err error
	if isSame := utils.VerifyPassword(model.Password, req.CurrentPassword); !isSame {
		return nil, err
	}

	// hashing password and overwrite the request
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}
	
	// overwrite existing password
	model.Password = hashedPassword
	
	// save new password
	return s.repository.Update(model)
}