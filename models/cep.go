package models

type CEPResponse struct {
	Rua    string `json:"rua" example:"Rua dos Eldar"`
	Bairro string `json:"bairro" example:"Rivendell"`
	Cidade string `json:"cidade" example:"Mirkwood"`
	Estado string `json:"estado" example:"Arnor"`
}

type CEPErrorResponse struct {
	Error        string `json:"error" example:"CEP inválido"`
	CepInformado string `json:"cep" example:"00000000"`
}

type CEPResponseExterno struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}
