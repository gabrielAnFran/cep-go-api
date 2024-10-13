package cep_controller

import (
	"cep-gin-clean-arch/mocks"
	"cep-gin-clean-arch/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestBuscarCEPSucesso(t *testing.T) {
	CEPRepositoryInterface := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "99150000"}}

	expectedCEP := models.CEPResponse{Estado: "RS", Cidade: "Marau", Bairro: "Frei Adelar", Rua: "Festivo"}
	CEPRepositoryInterface.On("Buscar", "99150000").Return(expectedCEP, nil)
	mockService.On("BuscarCEP", "99150000").Return(&expectedCEP, nil)

	webCEPHandler := CEPWebHandler{CEPRepository: CEPRepositoryInterface, BuscaCepExterno: mockService}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expectedCEP, responseBody)
}

func TestBuscarCEPInvalido(t *testing.T) {
	CEPRepositoryInterface := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "00000000"}}

	CEPRepositoryInterface.On("Buscar", "00000000").Return(models.CEPResponse{}, errors.New("CEP not found"))
	mockService.On("BuscarCEP", "00000000").Return(nil, errors.New("CEP inválido"))

	webCEPHandler := CEPWebHandler{CEPRepository: CEPRepositoryInterface, BuscaCepExterno: mockService}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPErrorResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, "CEP inválido", responseBody.Error)
}

func TestBuscarCEPNaoNumerico(t *testing.T) {
	CEPRepositoryInterface := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "000000qq"}}

	CEPRepositoryInterface.On("Buscar", "000000qq").Return(models.CEPResponse{}, errors.New("CEP deve conter apenas dígitos numéricos"))
	mockService.On("BuscarCEP", "000000qq").Return(nil, errors.New("CEP deve conter apenas dígitos numéricos"))

	webCEPHandler := CEPWebHandler{CEPRepository: CEPRepositoryInterface, BuscaCepExterno: mockService}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPErrorResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, "CEP deve conter apenas dígitos numéricos", responseBody.Error)
}

func TestBuscarCEPInvalidFormat(t *testing.T) {
	// Setup mocks
	CEPRepositoryInterface := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)

	// Create a response recorder
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	c.Params = gin.Params{{Key: "cep", Value: "1234-567"}}

	webCEPHandler := CEPWebHandler{CEPRepository: CEPRepositoryInterface, BuscaCepExterno: mockService}

	// Call the handler
	webCEPHandler.BuscarCEP(c)

	// Check response
	responseBody := models.CEPErrorResponse{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, "CEP deve conter apenas dígitos numéricos", responseBody.Error)
}

func TestCEPWebHandlerInitialization(t *testing.T) {
	buscarCEPRepository := new(mocks.CEPRepositoryInterface)
	buscaCepExterno := new(mocks.MockCEPService)

	handler := &CEPWebHandler{
		CEPRepository:   buscarCEPRepository,
		BuscaCepExterno: buscaCepExterno,
	}

	assert.Equal(t, buscarCEPRepository, handler.CEPRepository)
	assert.Equal(t, buscaCepExterno, handler.BuscaCepExterno)
}

func TestBuscarCEP_ErrorHandling(t *testing.T) {
	gin.SetMode(gin.TestMode)

	CEPRepositoryInterface := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{{Key: "cep", Value: "00000000"}}

	CEPRepositoryInterface.On("Buscar", "00000000").Return(models.CEPResponse{}, errors.New("CEP inválido"))
	mockService.On("BuscarCEP", "00000000").Return(nil, errors.New("CEP inválido"))

	handler := CEPWebHandler{CEPRepository: CEPRepositoryInterface, BuscaCepExterno: mockService}

	handler.BuscarCEP(c)

	responseBody := models.CEPErrorResponse{}
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "CEP inválido", responseBody.Error)
}

func TestNewBuscarCEPHandler(t *testing.T) {
	buscarCEPRepository := new(mocks.CEPRepositoryInterface)
	mockService := new(mocks.MockCEPService)

	handler := NewBuscarCEPHandler(buscarCEPRepository, mockService)

	assert.Equal(t, buscarCEPRepository, handler.CEPRepository)
	assert.Equal(t, mockService, handler.BuscaCepExterno)
}
