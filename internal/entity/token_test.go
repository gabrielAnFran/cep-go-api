package entity

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTokenLoginIsValidTokenLoginValid(t *testing.T) {
	token := &TokenLogin{Email: "test@example.com", Senha: "password"}

	err := token.IsValidTokenLogin()

	assert.Equal(t, nil, err)

}

func TestTokenLoginIsValidTokenLoginEmptyEmail(t *testing.T) {
	token := &TokenLogin{Email: "", Senha: "password"}
	errorMsg := "Email não pode ser vazio"

	err := token.IsValidTokenLogin()

	assert.NotEqual(t, nil, err)
	assert.Equal(t, errorMsg, err.Error())

}

func TestTokenLoginIsValidTokenLoginEmptySenha(t *testing.T) {
	token := &TokenLogin{Email: "test@example.com", Senha: ""}
	errorMsg := "Senha nao pode ser vazia"

	err := token.IsValidTokenLogin()

	assert.NotEqual(t, nil, err)
	assert.Equal(t, errorMsg, err.Error())

}

func TestTokenLoginIsValidTokenLoginInvalidEmailFormat(t *testing.T) {
	token := &TokenLogin{Email: "testexample.com", Senha: "password"}
	errorMsg := "Email inválido"

	err := token.IsValidTokenLogin()

	assert.NotEqual(t, nil, err)
	assert.Equal(t, errorMsg, err.Error())
	
}

func TestNewTokenLogin(t *testing.T) {
	email := "test@example.com"
	senha := "password"
	token := NewTokenLogin(email, senha)

	assert.Equal(t, email, token.Email)
	assert.Equal(t, senha, token.Senha)

}
