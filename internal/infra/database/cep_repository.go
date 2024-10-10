package database

import (
	"cep-gin-clean-arch/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var url = "http://viacep.com.br"

type CEPRepository struct{}

func NewCEPRepository() *CEPRepository {
	return &CEPRepository{}
}

type Endereco struct {
	CEP    string `json:"CEP"`
	Estado string `json:"Estado"`
	Cidade string `json:"Cidade"`
	Bairro string `json:"Bairro"`
	Rua    string `json:"Rua"`
}

type ViaCEPResponse struct {
	Uf         string `json:"uf"`
	Localidade string `json:"localidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}

func (r *CEPRepository) Buscar(cep string) (models.CEPResponse, error) {

	// Primeiro busca os dados "em memória"
	dados := []byte(jsonDados)

	var addresses []Endereco

	// Unmarshal JSON data
	err := json.Unmarshal(dados, &addresses)
	if err != nil {
		return models.CEPResponse{}, errors.New("Erro ao acessar dados de CEP")
	}

	// Create a map for quick lookup
	addressMap := make(map[string]Endereco)
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

	// Como uma ultima opção, busca o CEP na API externa
	viaCepResponse, err := buscarViaCep(cep)
	if err != nil {
		return models.CEPResponse{}, errors.New("CEP não encontrado")
	}

	return viaCepResponse, nil
}

func buscarViaCep(cep string) (models.CEPResponse, error) {
	URLcompleta := fmt.Sprintf("%s/ws/%s/json/", url, cep)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URLcompleta, nil)
	if err != nil {
		return models.CEPResponse{}, errors.New("CEP não encontrado")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.CEPResponse{}, errors.New("CEP não encontrado")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CEPResponse{}, errors.New("CEP não encontrado")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.CEPResponse{}, errors.New("CEP não encontrado")
	}

	var viaCepResponse ViaCEPResponse

	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		return models.CEPResponse{}, errors.New("CEP não encontrado")
	}

	return models.CEPResponse{
		Estado: viaCepResponse.Uf,
		Cidade: viaCepResponse.Localidade,
		Bairro: viaCepResponse.Bairro,
		Rua:    viaCepResponse.Logradouro,
	}, nil
}
