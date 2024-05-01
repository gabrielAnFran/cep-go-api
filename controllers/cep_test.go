package controllers

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
	c.Params = gin.Params{{Key: "cep", Value: "01001000"}}

	expectedCEP := models.CEPResponse{Estado: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"}
	serviceCEP.On("Buscar", "01001000").Return(expectedCEP, nil)

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
