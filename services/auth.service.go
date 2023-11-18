package services

import (
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/utils"
)

type AuthService struct {
	config *initializers.Config
	// repository *repositories.UserRepository
}

func NewAuthService(config *initializers.Config) *AuthService {
	return &AuthService{
		config: config,
	}
}

func (s AuthService) Register() (string, error) {
	// generate jwt token
	credential := &utils.UserCredential{
		ID: "asdasd-asdasd-adasds-asdasd",
		Name: "asdadasdas",
	}
	jwtToken, err := utils.GenerateJwtToken(s.config.JWTSECRET, credential, s.config.JWTEXPIRED);

	if err != nil {
		return "", err
	}
	
	return jwtToken, err
}	