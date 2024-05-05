package database

import (
	"cep-gin-clean-arch/models"
	"os"
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

func TestBuscarNaSupabaseHostInexistente(t *testing.T) {
	repo := NewCEPRepository()

	os.Setenv("SUPABASE_URL", "http://localhost:5432")
	os.Setenv("SUPABASE_KEY", "123456789")

	// CEP q só tem na supabase
	cep := "99150000"

	result, err := repo.Buscar(cep)

	assert.NotEmpty(t, err)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "connection refused")

}
