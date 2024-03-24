package services

import (
	"golang.org/x/crypto/bcrypt"
)

func (a Auth) HashPassword(password string) (string, error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashpass), nil
}

func (a Auth) ValidatePassword(password string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
