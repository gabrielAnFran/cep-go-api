package entity

import (
	"errors"
	"strconv"
)

type CEP struct {
	Cep string `json:"cep"`
}

// Camada que trata a regra de negócio da API
func NewCep(cep string) *CEP {

	return &CEP{
		Cep: cep,
	}
}

func (c *CEP) IsValidCep(cep string) error {

	if len(cep) == 0 {
		return errors.New("CEP não pode ser vazio")
	}

	_, err := strconv.Atoi(cep)
	if err != nil {
		return errors.New("CEP deve conter apenas dígitos numéricos")
	}

	if len(cep) != 8 {
		return errors.New("CEP inválido")
	}

	if cep == "00000000" {
		return errors.New("CEP inválido")
	}

	return nil
}
