package models

type CEPResponse struct {
	Rua    string `json:"rua" example:"Rua dos Eldar"`
	Bairro string `json:"bairro" example:"Rivendell"`
	Cidade string `json:"cidade" example:"Mirkwood"`
	Estado string `json:"estado" example:"Arnor"`
}

type CEPErrorResponse struct {
	Error        string `json:"error"`
	CepInformado string `json:"cep"`
}
