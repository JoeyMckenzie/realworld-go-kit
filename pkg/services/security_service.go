package services

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	SecurityService interface {
		HashPassword(rawPassword string) (string, error)
		PasswordIsValid(existingPassword, rawPassword string) bool
	}

	securityService struct{}
)

func NewSecurityService() SecurityService {
	return &securityService{}
}

func (s *securityService) HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *securityService) PasswordIsValid(existingPassword, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(rawPassword))
	return err == nil
}
