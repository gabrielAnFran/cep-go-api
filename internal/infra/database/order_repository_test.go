package database

import (
	"cep-gin-clean-arch/models"
	"testing"
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
