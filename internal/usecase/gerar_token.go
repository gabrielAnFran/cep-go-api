package usecase

import (
	"cep-gin-clean-arch/configs"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt"
)

type ServiceAuth struct{}

func (ServiceAuth) GenerateTokenJWT() (string, error) {
	now := time.Now().UTC()

	// Gerando um ID unico para o token
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	claims := jwt.MapClaims{
		"jti": uuid,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(time.Hour * 1).Unix(),
		"sub": 1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	configs, err := configs.LoadConfig("../../.")
	if err != nil {
		return "", errors.New("Erro ao carregar variáveis de ambiente")
	}

	spew.Dump(configs)
	tokenString, err := token.SignedString([]byte(configs.JWTSecret))
	if err != nil {
		return "", err
	}

	// Mock de um token para retorno

	return tokenString, nil
}
