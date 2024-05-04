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
	// Execute é a função responsável por buscar um CEP.
	// Se o CEP não for encontrado após adicionar os zeros, a função retorna um erro indicando que o CEP não foi encontrado.
	// Caso contrário, retorna os detalhes do endereço correspondente ao CEP encontrado.
	var cep string
	if input != nil {
		cep = *input
	}

	// Loop para adicionar zeros à direita do CEP.
	for i := 1; i < 8; i++ {
		cepResponse, err := b.CEPRepository.Buscar(cep) // Chama o método Buscar do repositório para obter informações do CEP.
		if err != nil {                                 // Verifica se houve um erro na busca do CEP.
			if err.Error() == "CEP não encontrado" { // Verifica se o erro indica que o CEP não foi encontrado.
				cep = cep[:8-i] + strings.Repeat("0", i) //  Caso seja, adiciona um zero à direita do CEP.
				continue                                 // Vai para mais uma iteração do loop com o CEP modificado.
			} else {
				return BuscarCepOutputDTO{}, err // Retorna um erro se ocorrer um erro diferente de "CEP não encontrado".
			}
		} else {
			return BuscarCepOutputDTO{Rua: cepResponse.Rua, Bairro: cepResponse.Bairro, Cidade: cepResponse.Cidade, Estado: cepResponse.Estado}, nil // Retorna os detalhes do endereço correspondente ao CEP encontrado.
		}
	}
	return BuscarCepOutputDTO{}, errors.New("CEP não encontrado") // Retorna um erro indicando que o CEP não foi encontrado após o loop.
}
