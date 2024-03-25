package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserCredential struct {
	ID string
}

func GenerateJwtToken(secret string, credential *UserCredential, expiredInSeconds int) (string, error) {
	var (
		key []byte
		t *jwt.Token
		s string
	)

	key = []byte(secret)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": credential.ID,
		"exp": time.Now().Add(time.Duration(expiredInSeconds) * time.Second).Unix(), //
	})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}