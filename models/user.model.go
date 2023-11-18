package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID *uuid.UUID `gorm:"column:uuid;type:char(36)"`
	Email string `gorm:"column:email;type:varchar;size:100"`
	Password string `gorm:"column:password;type:varchar;size:255"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar;size:20"`
	IsActive bool `gorm:"column:is_active;type:tinyint"`
}

type RegisterRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=1,max=8"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}