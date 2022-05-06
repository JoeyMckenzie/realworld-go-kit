package services

import "github.com/stretchr/testify/mock"

type MockSecurityService struct {
	mock.Mock
}

func NewMockSecurityService() SecurityService {
	return &MockSecurityService{}
}

func (mss *MockSecurityService) ResetMocks() {
	mss.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (m *MockSecurityService) HashPassword(rawPassword string) (string, error) {
	args := m.Called(rawPassword)
	return args.String(0), args.Error(1)
}

func (m *MockSecurityService) IsValidPassword(existingPassword, rawPassword string) bool {
	args := m.Called(existingPassword, rawPassword)
	return args.Bool(0)
}
