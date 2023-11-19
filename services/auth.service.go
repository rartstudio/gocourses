package services

import (
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/utils"
)

type AuthService struct {
	config *initializers.Config
	repository *repositories.UserRepository
	otp *OtpService
	jwt *JWTService
}

func NewAuthService(config *initializers.Config, repository *repositories.UserRepository, otp *OtpService, jwt *JWTService) *AuthService {
	return &AuthService{
		config: config,
		repository: repository,
		otp: otp,
		jwt: jwt,
	}
}

func (s AuthService) Register(req *models.RegisterRequest) (string, error) {
	// hashing password and overwrite the request
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}
	req.Password = hashedPassword

	// write it to model
	model := models.WriteToModelUser(req)

	// save it to database
	// s.repository.Create(model)

	// processing jwt
	jwtToken, err := s.jwt.ProcessingJwtToken(model)
	if err != nil {
		return "", err
	}

	// processing otp
	_, err = s.otp.ProcessingOtp(model)
	if err != nil {
		return "", err
	}
	
	return jwtToken, err
}	