package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID *uuid.UUID `gorm:"column:uuid;type:char(36)"`
	Email string `gorm:"column:email;type:varchar;size:100"`
	Password string `gorm:"column:password;type:varchar;size:255"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar;size:20"`
	VerifiedDate *time.Time `gorm:"column:verified_date;type:datetime"`
	IsActive bool `gorm:"column:is_active;type:tinyint"`
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

func WriteToModelUser(req *RegisterRequest) *User {
	model := &User{}

	// Generate a new UUID
	newUUID := uuid.New()

	model.Email = req.Email
	model.Password = req.Password
	model.PhoneNumber = req.PhoneNumber
	model.UUID = &newUUID

	return model
}

type UserResponse struct {
	UUID *uuid.UUID `json:"uuid"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func FilterUserResponseV1(model *User) UserResponse {
	return UserResponse{
		UUID: model.UUID,
		Email: model.Email,
		PhoneNumber: model.PhoneNumber,
	}
}