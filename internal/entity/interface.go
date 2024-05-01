package entity

import "cep-gin-clean-arch/models"

type CEPRepositoryInterface interface {
	Buscar(string) (models.CEPResponse, error)
}

type GerarTokenInterface interface {
	GenerateTokenJWT() (string, error)
}
