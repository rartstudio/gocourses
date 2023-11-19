package clients

import "gopkg.in/gomail.v2"

type EmailService struct {
	mail *gomail.Dialer
}

func NewEmailService(mail *gomail.Dialer) *EmailService {
	return &EmailService{mail: mail}
}

func (service EmailService) SendEmail(from string, to string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	err := service.mail.DialAndSend(m)
	if err != nil {
		panic(err)
	}

	return err
}