package services

import (
	"cep-gin-clean-arch/models"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// BuscaCepExternoService handles external CEP lookup
type BuscaCepExternoService struct{}

// NewBuscaCepExternoService creates a new instance of BuscaCepExternoService
func NewBuscaCepExternoService() *BuscaCepExternoService {
	return &BuscaCepExternoService{}
}

func (s *BuscaCepExternoService) BuscaCEP(cep string) (models.CEPResponse, error) {

	ctx := context.Background()

	url := "http://viacep.com.br/ws/" + cep + "/json/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.CEPResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.CEPResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.CEPResponse{}, err
	}
	spew.Dump(err)

	spew.Dump(string(body))

	var cepResponse models.CEPResponseExterno
	err = json.Unmarshal(body, &cepResponse)
	if err != nil {
		return models.CEPResponse{}, err
	}

	cepResponseExterno := models.CEPResponse{
		Rua:    cepResponse.Logradouro,
		Bairro: cepResponse.Bairro,
		Cidade: cepResponse.Localidade,
		Estado: cepResponse.UF,
	}

	return cepResponseExterno, nil
}
