package entity

import (
	"cep-gin-clean-arch/models"
)

type GerarTokenInterface interface {
	GenerateTokenJWT(models.TokenLoginRequest) (string, error)
}
