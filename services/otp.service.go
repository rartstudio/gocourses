package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/models"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

type OtpService struct {
	config *initializers.Config
	redisOtp *redis.Client
	mail *gomail.Dialer
}

func NewOtpService(config *initializers.Config, redisOtp *redis.Client, mail *gomail.Dialer) *OtpService {
	return &OtpService{config: config, redisOtp: redisOtp, mail: mail}
}


func (s OtpService) GenerateOTP(digits int) string {
	rand.Seed(time.Now().UnixNano())
	min := int64(1)
	max := int64(1)

	for i := 0; i < digits; i++ {
		max *= 10
	}

	randomNumber := rand.Int63n(max-min) + min
	return fmt.Sprintf("%0*d", digits, randomNumber)
}

func (s OtpService) ProcessingOtp(model *models.User) (string, error) {
	// generate otp
	otp := s.GenerateOTP(6)

	key := model.UUID.String() + "-otp"

	// storing otp
	err := s.redisOtp.Set(context.Background(), key, otp, 3*time.Minute).Err()
	if err != nil {
		fmt.Println("Error storing otp in redis", err)
		return "", err
	}

	// sending otp to email
	m := gomail.NewMessage()
	m.SetHeader("From", "admin222@gmail.com")
	m.SetHeader("To", model.Email)
	message := fmt.Sprintf("Hello here is your %v", otp)
	m.SetBody("text/html", message)
	err = s.mail.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return otp, nil
}