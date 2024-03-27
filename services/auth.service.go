package services

import (
	"errors"

	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/utils"
)

type AuthService struct {
	config *initializers.Config
	otp *OtpService
	jwt *JWTService
	userService *UserService
}

func NewAuthService(config *initializers.Config, repository *repositories.UserRepository, otp *OtpService, jwt *JWTService, userService *UserService) *AuthService {
	return &AuthService{
		config: config,
		otp: otp,
		jwt: jwt,
		userService: userService,
	}
}

func (s AuthService) Register(req *models.RegisterRequest) (string, error) {
	// hashing password and overwrite the request
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", errors.New("gagal hash password")
	}
	req.Password = hashedPassword

	// write it to model
	model := models.WriteToModelUser(req)

	// save it to database
	model, err = s.userService.Create(model);
	if err != nil {
		return "", err
	}

	// processing jwt
	jwtToken, err := s.jwt.ProcessingJwtToken(model)
	if err != nil {
		return "", errors.New("gagal mendapatkan token")
	}

	// processing otp
	_, err = s.otp.ProcessingOtp(model)
	if err != nil {
		return "", errors.New("gagal mendapatkan otp")
	}
	
	return jwtToken, err
}	

func (s AuthService) Login(req *models.LoginRequest, model *models.User) (string, error) {
	var err error
	if isSame := utils.VerifyPassword(model.Password, req.Password); !isSame {
		return "", errors.New("password tidak sama")
	}

	// processing jwt
	jwtToken, err := s.jwt.ProcessingJwtToken(model)
	if err != nil {
		return "", errors.New("gagal mendapatkan token")
	}

	return jwtToken, err
}

func (s AuthService) VerifyAccount(req *models.VerifyAccountRequest, model *models.User) (*models.User, error) {
	otp, err := s.otp.GetOtp(model)
	if err != nil {
		return nil, err
	}

	// checking otp from redis same or not 
	if otp != req.Otp {
		return nil, err
	}

	// update user data is active
	model = models.WriteToModelUserIsActive(model)
	model, err = s.userService.UpdateUser(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}