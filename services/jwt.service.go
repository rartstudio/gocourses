package services

import (
	"context"
	"fmt"
	"time"

	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/utils"
	"github.com/redis/go-redis/v9"
)

type JWTService struct {
	config *initializers.Config
	redisToken *redis.Client
}

func NewJWTService(config *initializers.Config, redisToken *redis.Client) *JWTService {
	return &JWTService{config: config, redisToken: redisToken}
}

func (s JWTService) ProcessingJwtToken(model *models.User) (string, error) {
	// generate jwt token
	credential := &utils.UserCredential{
		ID: model.UUID.String(),
	}
	jwtToken, err := utils.GenerateJwtToken(s.config.JWTSECRET, credential, s.config.JWTEXPIRED);

	if err != nil {
		return "", err
	}

	// saving it to redis	
	key := model.UUID.String() + "-jwt"
	err = s.redisToken.Set(context.Background(), key, jwtToken, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error storing token in redis", err)
		return "", err
	}

	return jwtToken, nil
}

func (s JWTService) RetrieveJwtTokenFromRedis(uuid string)  (string, error) {
	key := uuid + "-jwt"

	// retrieve otp
	r := s.redisToken.Get(context.Background(), key)
	if r.Err() != nil {
		fmt.Println("Error get jwt from redis", r.Err())
		return "", r.Err()
	}

	return r.Val(), nil
}