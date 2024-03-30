package services

import (
	"context"
	"encoding/json"
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

type OtpMessageEmail struct {
	Otp string `json:"otp"`
	Email string `json:"email"`
}

func NewOtpService(config *initializers.Config, redisOtp *redis.Client, mail *gomail.Dialer) *OtpService {
	return &OtpService{config: config, redisOtp: redisOtp, mail: mail}
}

func (s OtpService) GenerateOTP(digits int) string {
	min := int64(1)
	max := int64(1)

	for i := 0; i < digits; i++ {
		max *= 10
	}

	randomNumber := rand.Int63n(max-min) + min
	return fmt.Sprintf("%0*d", digits, randomNumber)
}

func (s OtpService) RemoveOtpFromRedis(model *models.User) (string, error) {
	key := model.UUID.String() + "-otp"

	// remove otp
	r := s.redisOtp.Del(context.Background(), key)
	if r.Err() != nil {
		fmt.Println("Error remove otp from redis", r.Err())
		return "", r.Err()
	}

	return "", nil
}

func (s OtpService) RetrieveOtpFromRedis(model *models.User) (string, error) {
	key := model.UUID.String() + "-otp"

	// retrieve otp
	r := s.redisOtp.Get(context.Background(), key)
	if r.Err() != nil {
		fmt.Println("Error get otp from redis", r.Err())
		return "", r.Err()
	}

	return r.Val(), nil
}

func (s OtpService) GetOtp(model *models.User) (string, error) {
	// generate otp
	otp := s.GenerateOTP(6)

	key := model.UUID.String() + "-otp"

	// storing otp
	err := s.redisOtp.Set(context.Background(), key, otp, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Error storing otp in redis", err)
		return "", err
	}

	// message to pass queue
	message := OtpMessageEmail{
		Otp: otp,
		Email: model.Email,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error json encode")
	}

	// publish the OTP event
	err = s.redisOtp.Publish(context.Background(), "otp_channel", messageBytes).Err()
	if err != nil {
		fmt.Println("Error publish message")
	}
	return otp, nil
}

func (s OtpService) HandleEmails() {
	pubsub := s.redisOtp.Subscribe(context.Background(), "otp_channel")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			fmt.Println("Error receiving message from redis ")
			continue
		}

		var message OtpMessageEmail
		err = json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			fmt.Println("Error decoding json")
		}

		go s.sendEmail(message)
	}
}

func (s OtpService) sendEmail(message OtpMessageEmail) {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin222@gmail.com")
	m.SetHeader("To", message.Email)
	msg := fmt.Sprintf("Hello here is your %v", message.Otp)
	m.SetBody("text/html", msg)
	err := s.mail.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}