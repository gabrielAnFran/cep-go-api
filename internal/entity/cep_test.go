// internal/entity/cep_test.go

package entity

import (
	"testing"
)

func TestIsValidCepValido(t *testing.T) {
	cep := "12345678"
	c := &CEP{}

	err := c.IsValidCep(cep)

	if err != nil {
		t.Errorf("Nenhum erro esperado, recebeu %v", err)
	}
}

func TestIsValidCepVazio(t *testing.T) {
	cep := ""
	c := &CEP{}

	err := c.IsValidCep(cep)

	if err == nil {
		t.Error("Esperado erro para CEP vazio")
	}

	if err.Error() != "CEP não pode ser vazio" {
		t.Errorf("Erro esperado CEP deve conter apenas dígitos numéricos, recebeu %v", err)
	}
}

func TestIsValidCepInvalido(t *testing.T) {
	cep := "00000000"
	c := &CEP{}

	err := c.IsValidCep(cep)

	if err == nil {
		t.Error("Erro esperado e não retornado para CEP")
	}

	if err.Error() != "CEP inválido" {
		t.Errorf("Erro esperado CEP deve conter apenas dígitos numéricos, recebeu %v", err)
	}
}

func TestIsValidCepLetras(t *testing.T) {
	cep := "000000a@"
	c := &CEP{}

	err := c.IsValidCep(cep)

	if err == nil {
		t.Error("Erro esperado e não retornado para CEP com letras.")
	}

	if err.Error() != "CEP deve conter apenas dígitos numéricos" {
		t.Errorf("Erro esperado CEP deve conter apenas dígitos numéricos, recebeu %v", err)
	}
}
