package initializers

import (
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"
)

func ConnectToEmailProvider(config *Config) (*gomail.Dialer) {
	d := gomail.NewDialer(config.EMAILHOST, config.EMAILPORT, config.EMAILUSERNAME, config.EMAILPASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	log.Println("🚀 Connected Successfully to the Email Provider")

	return d
}