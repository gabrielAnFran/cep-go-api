package database

import (
	"cep-gin-clean-arch/models"
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
