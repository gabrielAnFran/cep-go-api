package usecase

import (
	"cep-gin-clean-arch/mocks"
	"cep-gin-clean-arch/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuscarCEPUseCase(t *testing.T) {
	mockRepo := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)
	useCase := NewBuscarCEPUseCase(mockRepo, mockService)
	assert.NotNil(t, useCase)
}

func TestExecute(t *testing.T) {
	mockRepo := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)
	useCase := NewBuscarCEPUseCase(mockRepo, mockService)

	mockRepo.
		On("Buscar", "11111111").Return(models.CEPResponse{
		Estado: "RJ", Cidade: "Rio de Janeiro", Bairro: "Inhaúma", Rua: "Rua José dos Reis",
	}, nil).
		On("Buscar", "00000000").Return(models.CEPResponse{}, nil)

	cep := "11111111"
	useCaseCep := BuscarCEPuseCase{useCase.CEPRepository, useCase.BuscaCepExterno}
	res, err := useCaseCep.CEPRepository.Buscar(cep)
	assert.Equal(t, nil, err)
	assert.Equal(t, models.CEPResponse{
		Estado: "RJ", Cidade: "Rio de Janeiro", Bairro: "Inhaúma", Rua: "Rua José dos Reis",
	}, res)
}

func TestBuscarCEPuseCase_Execute_SuccessfullyFindCEPInRepository(t *testing.T) {
	mockRepo := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)
	useCase := BuscarCEPuseCase{mockRepo, mockService}

	cep := "11111111"
	expectedResponse := models.CEPResponse{
		Estado: "RJ", Cidade: "Rio de Janeiro", Bairro: "Inhaúma", Rua: "Rua José dos Reis",
	}

	mockRepo.On("Buscar", cep).Return(expectedResponse, nil)

	input := cep
	output, err := useCase.Execute(&input)

	assert.NoError(t, err)
	assert.Equal(t, BuscarCepOutputDTO{
		Rua: expectedResponse.Rua, Bairro: expectedResponse.Bairro, Cidade: expectedResponse.Cidade, Estado: expectedResponse.Estado,
	}, output)
}

func TestBuscarCEPuseCase_Execute_ReturnErrorWhenNotFound(t *testing.T) {
	mockRepo := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)
	useCase := BuscarCEPuseCase{mockRepo, mockService}

	cep := "30000000"

	mockRepo.On("Buscar", cep).Return(models.CEPResponse{}, errors.New("CEP não encontrado"))
	mockService.On("BuscaCEP", cep).Return(models.CEPResponse{}, errors.New("CEP não encontrado"))

	input := cep
	_, err := useCase.Execute(&input)

	assert.Error(t, err)
	assert.Equal(t, "CEP não encontrado", err.Error())
}
