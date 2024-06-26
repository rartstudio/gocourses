package repositories

import (
	"github.com/rartstudio/gocourses/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r UserRepository) Create(model *models.User) (*models.User, error){
	return model, r.DB.Create(model).Error
}

func (r UserRepository) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.DB.Take(&user, "email = ?", email).Error
	return user, err
}

func (r UserRepository) GetByUuid(uuid string) (*models.User, error) {
	var user *models.User
	err := r.DB.Preload("Profile").Take(&user, "uuid = ?", uuid).Error
	return user, err
}

func (r UserRepository) Update(model *models.User) (*models.User, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r UserRepository) AddUserProfile(model *models.UserProfile) (*models.UserProfile, error) {
	return model, r.DB.Create(model).Error
}

func (r UserRepository) UpdateUserProfile(model *models.UserProfile) (*models.UserProfile, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}