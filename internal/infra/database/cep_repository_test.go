package database

import (
	"cep-gin-clean-arch/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscarSucesso(t *testing.T) {
	repo := NewCEPRepository()

	cep := "01001000"

	expectedCEP := models.CEPResponse{Estado: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"}

	result, err := repo.Buscar(cep)

	assert.Empty(t, err)
	assert.Equal(t, expectedCEP, result)

}

func TestBuscarViaCep(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/ws/12345678/json/", r.URL.Path)

		mockResponse := map[string]string{
			"uf":         "SP",
			"localidade": "São Paulo",
			"bairro":     "Centro",
			"logradouro": "Avenida Paulista",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	originalURL := "http://viacep.com.br"
	url = server.URL

	result, err := buscarViaCep("12345678")

	url = originalURL

	assert.NoError(t, err)

	expectedResult := models.CEPResponse{
		Estado: "SP",
		Cidade: "São Paulo",
		Bairro: "Centro",
		Rua:    "Avenida Paulista",
	}
	assert.Equal(t, expectedResult, result)
}
