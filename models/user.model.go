package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID *uuid.UUID `gorm:"column:uuid;type:char(36)"`
	Email string `gorm:"column:email;type:varchar(100)"`
	Password string `gorm:"column:password;type:varchar(255)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20)"`
	VerifiedDate *time.Time `gorm:"column:verified_date;type:datetime"`
	IsActive bool `gorm:"column:is_active;type:tinyint"`
}

type UserProfile struct {
	gorm.Model
	UserId uint `gorm:"column:user_id"`
	Name string `gorm:"column:name;type:varchar(255)"`
	ProfileImage string `gorm:"column:profile_image;type:varchar(255)"`
	User *User `gorm:"foreignKey:user_id;references:id"`
}

type UserAccount struct {
	gorm.Model
	UserId uint `gorm:"column:user_id"`
	AccountNumber uint `gorm:"column:account_number;type:int(12)"`
	Balance uint `gorm:"column:balance;type:int"`
	User *User `gorm:"foreignKey:user_id;references:id"`
}

type RegisterRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=1,max=8"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=1,max=8"`
}

type VerifyAccountRequest struct {
	Otp string `json:"otp" validate:"required"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserProfileRequest struct {
	Name string `json:"name"`
	ProfileImage string `json:"profile_image"`
}

func WriteToModelUser(req *RegisterRequest) *User {
	model := &User{}

	// Generate a new UUID
	newUUID := uuid.New()

	model.Email = req.Email
	model.Password = req.Password
	model.PhoneNumber = req.PhoneNumber
	model.IsActive = true
	model.UUID = &newUUID

	return model
}

func WriteToModelUserProfile(req *UserProfileRequest, user *User) *UserProfile {
	model := &UserProfile{}

	model.Name = req.Name
	model.ProfileImage = req.ProfileImage
	model.UserId = user.ID

	return model
}

func WriteToModelUserIsActive(user *User) *User {
	user.IsActive = true;

	return user
}

type UserResponse struct {
	UUID *uuid.UUID `json:"uuid"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserProfileResponse struct {
	Name string `json:"name"`
	ProfileImage string `json:"profile_image"`
	User UserResponse `json:"user"`
}

func FilterUserResponseV1(model *User) UserResponse {
	return UserResponse{
		UUID: model.UUID,
		Email: model.Email,
		PhoneNumber: model.PhoneNumber,
	}
}

func FilterUserProfileResponseV1(model *UserProfile, user *User) UserProfileResponse {
	userResponse := FilterUserResponseV1(user);

	userProfileResponse := UserProfileResponse{
		Name: model.Name,
		ProfileImage: model.ProfileImage,
		User: userResponse,
	}

	return userProfileResponse
}