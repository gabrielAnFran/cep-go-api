package mocks

import (
	models "cep-gin-clean-arch/models"

	mock "github.com/stretchr/testify/mock"
)

// MockCEPService is a mock implementation of entity.CEPServiceInterface
type MockCEPService struct {
	mock.Mock
}

func (m *MockCEPService) BuscaCEP(cep string) (models.CEPResponse, error) {
	args := m.Called(cep)

	// Check if the first return value is nil or of type *entity.CEP
	if args.Get(0) == nil {
		return models.CEPResponse{}, args.Error(1)
	}

	return models.CEPResponse{}, args.Error(1)
}
