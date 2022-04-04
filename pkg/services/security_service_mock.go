package services

import "github.com/stretchr/testify/mock"

type MockSecurityService struct {
	mock.Mock
}

func NewMockSecurityService() *MockSecurityService {
	return &MockSecurityService{}
}

func (m *MockSecurityService) HashPassword(rawPassword string) (string, error) {
	args := m.Called(rawPassword)
	return args.String(0), args.Error(1)
}

func (m *MockSecurityService) PasswordIsValid(existingPassword, rawPassword string) bool {
	args := m.Called(existingPassword, rawPassword)
	return args.Bool(0)
}
