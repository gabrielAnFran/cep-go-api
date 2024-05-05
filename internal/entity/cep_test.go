// internal/entity/cep_test.go

package entity

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsValidCepValido(t *testing.T) {
	cep := "12345678"
	c := &CEP{}

	err := c.IsValidCep(cep)

	assert.Equal(t, nil, err)
}

func TestIsValidCepVazio(t *testing.T) {
	cep := ""
	c := &CEP{}

	err := c.IsValidCep(cep)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "CEP não pode ser vazio", err.Error())

}

func TestIsValidCepInvalido(t *testing.T) {
	cep := "00000000"
	c := &CEP{}

	err := c.IsValidCep(cep)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "CEP inválido", err.Error())
}

func TestIsValidCepLetras(t *testing.T) {
	cep := "000000a@"
	c := &CEP{}

	err := c.IsValidCep(cep)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "CEP deve conter apenas dígitos numéricos", err.Error())

}

func TestIsValidTamanhoInvalido(t *testing.T) {
	cep := "1111111111"
	c := &CEP{}

	err := c.IsValidCep(cep)

	assert.NotEqual(t, nil, err)
	assert.Equal(t,"CEP deve conter 8 dígitos numéricos", err.Error())
	
}
