package entity

import (
	"errors"
	"strings"
)

type TokenLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func NewTokenLogin(email string) *TokenLogin {
	return &TokenLogin{
		Email: email,
	}
}

func (t *TokenLogin) IsValidTokenLogin() error {

	if len(t.Email) == 0 {
		return errors.New("Email não pode ser vazio")
	}

	if len(t.Senha) == 0 {
		return errors.New("Senha nao pode ser vazia")
	}

	if !strings.Contains(t.Email, "@") || !strings.Contains(t.Email, ".") {
		return errors.New("Email inválido")
	}
	return nil

}
