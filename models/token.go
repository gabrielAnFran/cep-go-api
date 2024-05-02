package models

type TokenErrorResponse struct {
	Error string `json:"error" example:"Ocorreu um erro ao gerar o token"`
}

type TokenLoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
