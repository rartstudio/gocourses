package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	hashedPassword := string(salt)

	return hashedPassword, err
}

func VerifyPassword(hashedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))

	return err == nil
}