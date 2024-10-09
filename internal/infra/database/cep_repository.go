package database

import (
	"cep-gin-clean-arch/models"
	"encoding/json"
	"errors"
)

type CEPRepository struct{}

func NewCEPRepository() *CEPRepository {
	return &CEPRepository{}
}

type Address struct {
	CEP    string `json:"CEP"`
	Estado string `json:"Estado"`
	Cidade string `json:"Cidade"`
	Bairro string `json:"Bairro"`
	Rua    string `json:"Rua"`
}

func (r *CEPRepository) Buscar(cep string) (models.CEPResponse, error) {

	// Primeiro busca os dados "em memória"
	dados := []byte(jsonDados)

	var addresses []Address

	// Unmarshal JSON data
	err := json.Unmarshal(dados, &addresses)
	if err != nil {
		return models.CEPResponse{}, errors.New("Erro ao acessar dados de CEP")
	}

	// Create a map for quick lookup
	addressMap := make(map[string]Address)
	for _, address := range addresses {
		addressMap[address.CEP] = address
	}
	// Se o CEP estiver na memória, retorna os dados
	desiredAddress, found := addressMap[cep]
	if found {
		return models.CEPResponse{
			Estado: desiredAddress.Estado,
			Cidade: desiredAddress.Cidade,
			Bairro: desiredAddress.Bairro,
			Rua:    desiredAddress.Rua,
		}, nil
	}

	// Se não encontrar, retorna um erro
	return models.CEPResponse{}, errors.New("CEP não encontrado")

}
