package usecase

import (
	"cep-gin-clean-arch/models"
	"errors"
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

	t.Run("Valid CEP", func(t *testing.T) {
		cep := "12345678"
		mockRepo.On("Buscar", cep).Return(models.CEPResponse{Rua: "Test Street", Bairro: "Test District", Cidade: "Test City", Estado: "Test State"}, nil)
		input := &cep
		result, err := useCase.Execute(input)
		assert.Nil(t, err)
		assert.Equal(t, "Test Street", result.Rua)
		assert.Equal(t, "Test District", result.Bairro)
		assert.Equal(t, "Test City", result.Cidade)
		assert.Equal(t, "Test State", result.Estado)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CEP not found", func(t *testing.T) {
		cep := "00000000"
		mockRepo.On("Buscar", cep).Return(models.CEPResponse{}, errors.New("CEP não encontrado"))
		input := &cep
		_, err := useCase.Execute(input)
		assert.NotNil(t, err)
		assert.Equal(t, "CEP inválido", err.Error())
	})

	t.Run("Invalid CEP", func(t *testing.T) {
		cep := "invalid"
		input := &cep
		result, err := useCase.Execute(input)
		assert.NotNil(t, err)
		assert.Equal(t, "CEP inválido", err.Error())
		assert.Equal(t, BuscarCepOutputDTO{}, result)
	})
}
