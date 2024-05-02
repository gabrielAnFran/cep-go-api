package usecase

import (
	"cep-gin-clean-arch/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para CEPRepositoryInterface
type MockCEPRepository struct {
	mock.Mock
}

func (m *MockCEPRepository) Buscar(cep string) (models.CEPResponse, error) {
	args := m.Called(cep)
	return args.Get(0).(models.CEPResponse), args.Error(1)
}

func TestNewBuscarCEPUseCase(t *testing.T) {
	mockRepo := new(MockCEPRepository)
	useCase := NewBuscarCEPUseCase(mockRepo)
	assert.NotNil(t, useCase)
}

func TestExecute(t *testing.T) {
	mockRepo := new(MockCEPRepository)
	useCase := NewBuscarCEPUseCase(mockRepo)

	mockRepo.
		On("Buscar", "11111111").Return(models.CEPResponse{
		Estado: "RJ", Cidade: "Rio de Janeiro", Bairro: "Inhaúma", Rua: "Rua José dos Reis",
	}, nil).
		On("Buscar", "00000000").Return(models.CEPResponse{}, nil)

	cep := "11111111"
	useCaseCep := BuscarCEPuseCase{useCase.CEPRepository}
	res, err := useCaseCep.CEPRepository.Buscar(cep)
	assert.Equal(t, nil, err)
	assert.Equal(t, models.CEPResponse{
		Estado: "RJ", Cidade: "Rio de Janeiro", Bairro: "Inhaúma", Rua: "Rua José dos Reis",
	}, res)
}
