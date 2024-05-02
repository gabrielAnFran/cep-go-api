package usecase

import (
	"cep-gin-clean-arch/internal/entity"
	"errors"
	"strings"
)

type BuscarCepOutputDTO struct {
	Rua    string `json:"rua"    example:"Rua José dos Reis"`
	Bairro string `json:"bairro" example:"Inhaúma"`
	Cidade string `json:"cidade" example:"Rio de Janeiro"`
	Estado string `json:"estado" example:"RJ"`
}

type BuscarCEPuseCase struct {
	CEPRepository entity.CEPRepositoryInterface
}

func NewBuscarCEPUseCase(buscarCEPRepository entity.CEPRepositoryInterface) *BuscarCEPuseCase {
	return &BuscarCEPuseCase{
		CEPRepository: buscarCEPRepository,
	}
}

func (b *BuscarCEPuseCase) Execute(input *string) (BuscarCepOutputDTO, error) { //return BuscarCepOutputDTO{}, errors.New("Testinho")
	cep := entity.NewCep(*input)
	if err := cep.IsValidCep(*input); err != nil {
		return BuscarCepOutputDTO{}, err
	}

	// Lógica do loop para add zero a direita até encontrar o CEP

	for i := 1; i < 8; i++ {
		cepResponse, err := b.CEPRepository.Buscar(cep.Cep)
		if err != nil {
			if err.Error() == "CEP não encontrado" {
				cep.Cep = cep.Cep[:8-i] + strings.Repeat("0", i)
				continue // Continue the loop with the modified CEP
			} else {
				return BuscarCepOutputDTO{}, err
			}
		} else {
			return BuscarCepOutputDTO{Rua: cepResponse.Rua, Bairro: cepResponse.Bairro, Cidade: cepResponse.Cidade, Estado: cepResponse.Estado}, nil
		}
	}
	return BuscarCepOutputDTO{}, errors.New("CEP não encontrado")

}
