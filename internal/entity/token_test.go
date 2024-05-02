package entity

import (
	"testing"
)

func TestTokenLoginIsValidTokenLoginValid(t *testing.T) {
	token := &TokenLogin{Email: "test@example.com", Senha: "password"}
	wantErr := false
	errorMsg := ""

	err := token.IsValidTokenLogin()

	if (err != nil) != wantErr {
		t.Errorf("Erro de validação do token error = %v, wantErr %v", err, wantErr)
		return
	}
	if err != nil && err.Error() != errorMsg {
		t.Errorf("Erro de validação do token error message = %v, want %v", err.Error(), errorMsg)
	}
}

func TestTokenLoginIsValidTokenLoginEmptyEmail(t *testing.T) {
	token := &TokenLogin{Email: "", Senha: "password"}
	wantErr := true
	errorMsg := "Email não pode ser vazio"

	err := token.IsValidTokenLogin()

	if (err != nil) != wantErr {
		t.Errorf("Erro de validação do token error = %v, wantErr %v", err, wantErr)
		return
	}
	if err != nil && err.Error() != errorMsg {
		t.Errorf("Erro de validação do token error message = %v, want %v", err.Error(), errorMsg)
	}
}

func TestTokenLoginIsValidTokenLoginEmptySenha(t *testing.T) {
	token := &TokenLogin{Email: "test@example.com", Senha: ""}
	wantErr := true
	errorMsg := "Senha nao pode ser vazia"

	err := token.IsValidTokenLogin()

	if (err != nil) != wantErr {
		t.Errorf("Erro de validação do token error = %v, wantErr %v", err, wantErr)
		return
	}
	if err != nil && err.Error() != errorMsg {
		t.Errorf("Erro de validação do token error message = %v, want %v", err.Error(), errorMsg)
	}
}

func TestTokenLoginIsValidTokenLoginInvalidEmailFormat(t *testing.T) {
	token := &TokenLogin{Email: "testexample.com", Senha: "password"}
	wantErr := true
	errorMsg := "Email inválido"

	err := token.IsValidTokenLogin()

	if (err != nil) != wantErr {
		t.Errorf("Erro de validação do token error = %v, wantErr %v", err, wantErr)
		return
	}
	if err != nil && err.Error() != errorMsg {
		t.Errorf("Erro de validação do token error message = %v, want %v", err.Error(), errorMsg)
	}
}
