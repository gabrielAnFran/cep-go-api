package database

import (
	mockdb "cep-gin-clean-arch/mock-db"
	"cep-gin-clean-arch/models"
	"errors"
)

type CEPRepository struct{}

func NewCEPRepository() *CEPRepository {
	return &CEPRepository{}
}

func (r *CEPRepository) Buscar(cep string) (models.CEPResponse, error) {
	informacoesCEP, existe := mockdb.CEPS[cep]
	if existe {
		return informacoesCEP, nil
	}
	return models.CEPResponse{}, errors.New("CEP não encontrado")
}
