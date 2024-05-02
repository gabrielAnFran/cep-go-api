package usecase

import (
	"cep-gin-clean-arch/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenJWTValido(t *testing.T) {
	UsecaseAuth := UsecaseAuth{}

	token, err := UsecaseAuth.GenerateTokenJWT(models.TokenLoginRequest{
		Email: "meuemail@email.com",
		Senha: "minhasenha",
	})

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, token)
}
