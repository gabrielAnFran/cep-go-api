package database

import (
	"cep-gin-clean-arch/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscar(t *testing.T) {
	repo := NewCEPRepository()

	// Test case 1: CEP found
	cep := "01001000"

	expectedCEP := models.CEPResponse{Estado: "SP", Cidade: "São Paulo", Bairro: "Sé", Rua: "Praça da Sé"}

	result, err := repo.Buscar(cep)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expectedCEP {
		t.Errorf("Expected %v, got %v", expectedCEP, result)
	}

	// Test case 2: CEP not found
	cepNotFound := "00000000"
	_, err = repo.Buscar(cepNotFound)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
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
	//assert.Equal(t, "RS", result.Estado)

}
