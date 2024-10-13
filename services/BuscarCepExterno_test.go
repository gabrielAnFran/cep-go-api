package services

import (
	"cep-gin-clean-arch/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscaCEP(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP"
		}`))
	}))
	defer mockServer.Close()

	service := &BuscaCepExternoService{}

	cep := "01001-000"
	expectedResponse := models.CEPResponse{
		Rua:    "Praça da Sé",
		Bairro: "Sé",
		Cidade: "São Paulo",
		Estado: "SP",
	}

	response, err := service.BuscaCEP(cep)
	assert.NoError(t, err, "expected no error")

	assert.Equal(t, expectedResponse, response, "response should match expected response")
}
func TestNewBuscaCepExternoService(t *testing.T) {
	service := NewBuscaCepExternoService()
	assert.NotNil(t, service, "Expected non-nil instance of BuscaCepExternoService")
}
