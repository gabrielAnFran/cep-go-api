package cep_controller

import (
	"cep-gin-clean-arch/mocks"
	"cep-gin-clean-arch/models"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestBuscarCEPSucesso(t *testing.T) {
	serviceCEP := new(mocks.CEPRepositoryInterface)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "99150000"}}

	expectedCEP := models.CEPResponse{Estado: "RS", Cidade: "Marau", Bairro: "Frei Adelar", Rua: "Festivo"}
	serviceCEP.On("Buscar", "99150000").Return(expectedCEP, nil)

	webCEPHandler := CEPWebHandler{CEPRepository: serviceCEP}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	expectedResponseCEP := models.CEPResponse{}
	expectedResponseCEP.Estado = expectedCEP.Estado
	expectedResponseCEP.Cidade = expectedCEP.Cidade
	expectedResponseCEP.Bairro = expectedCEP.Bairro
	expectedResponseCEP.Rua = expectedCEP.Rua

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expectedResponseCEP, responseBody)

}

func TestBuscarCEPSucessoComCEPQueNaoExisteNoBanco(t *testing.T) {
	serviceCEP := new(mocks.CEPRepositoryInterface)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	// CEP que existe no banco
	// {"CEP": "95010000", "Estado": "RS", "Cidade": "Caxias do Sul", "Bairro": "Centro", "Rua": "Rua Sinimbu"},
	c.Params = gin.Params{{Key: "cep", Value: "95010111"}}
	expectedCEP := models.CEPResponse{Estado: "RS", Cidade: "Caxias do Sul", Bairro: "Centro", Rua: "Rua Sinimbu"}
	serviceCEP.On("Buscar", "95010111").Return(expectedCEP, nil)

	webCEPHandler := CEPWebHandler{CEPRepository: serviceCEP}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	expectedResponseCEP := models.CEPResponse{}
	expectedResponseCEP.Estado = expectedCEP.Estado
	expectedResponseCEP.Cidade = expectedCEP.Cidade
	expectedResponseCEP.Bairro = expectedCEP.Bairro
	expectedResponseCEP.Rua = expectedCEP.Rua

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expectedResponseCEP, responseBody)

}

func TestBuscarCEPInvalido(t *testing.T) {
	serviceCEP := new(mocks.CEPRepositoryInterface)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "00000000"}}

	serviceCEP.On("Buscar", "00000000").Return(models.CEPErrorResponse{Error: "CEP inválido"}, nil)

	webCEPHandler := CEPWebHandler{CEPRepository: serviceCEP}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPErrorResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, "CEP inválido", responseBody.Error)
}

func TestBuscarCEPNaoNumerico(t *testing.T) {
	serviceCEP := new(mocks.CEPRepositoryInterface)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Params = gin.Params{{Key: "cep", Value: "000000qq"}}

	serviceCEP.On("Buscar", "000000qq").Return(models.CEPErrorResponse{Error: "CEP deve conter apenas dígitos numéricos"}, nil)

	webCEPHandler := CEPWebHandler{CEPRepository: serviceCEP}

	webCEPHandler.BuscarCEP(c)

	responseBody := models.CEPErrorResponse{}

	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, nil, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, "CEP deve conter apenas dígitos numéricos", responseBody.Error)
}

func TestNewBuscarCEPHandler(t *testing.T) {
	mockRepository := new(mocks.CEPRepositoryInterface)

	handler := NewBuscarCEPHandler(mockRepository)

	if handler.CEPRepository != mockRepository {
		t.Errorf("Expected CEPRepository to be set to the mock repository")
	}
}
