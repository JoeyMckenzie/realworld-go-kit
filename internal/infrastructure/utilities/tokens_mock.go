package utilities

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) ResetMocks() {
	m.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (m *MockTokenService) GenerateUserToken(id uuid.UUID, email string) (string, error) {
	args := m.Called(id, email)
	return args.String(0), args.Error(1)
}
