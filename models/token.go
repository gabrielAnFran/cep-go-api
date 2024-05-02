package models

type TokenErrorResponse struct {
	Error string `json:"error" example:"Ocorreu um erro ao gerar o token"`
}

type TokenLoginRequest struct {
	Email string `json:"email" example:"email@email.com"`
	Senha string `json:"senha" example:"qualquerumamenosabre123"`
}

type TokenLoginResponse struct {
	Token string `json:"token" example:"eyJhbGcSSDSDiOiJIUzI1NiIsInRDSDF5cCI6DDSDADIkpXVCJ9SS.D.7Q6`
}
