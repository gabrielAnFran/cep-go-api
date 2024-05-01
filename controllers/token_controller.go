package controllers

import (
	"cep-gin-clean-arch/internal/entity"

	"github.com/gin-gonic/gin"
)

const (
	erroAoGerarToken = "Erro interno ao tentar gerar o token JWT"
)

type GerarTokenHandler struct {
	GerarTokenInterface entity.GerarTokenInterface
}

func NewGerarTokenHandler(gerarTokenRepository entity.GerarTokenInterface) *GerarTokenHandler {
	return &GerarTokenHandler{
		GerarTokenInterface: gerarTokenRepository,
	}
}

// @Tags Auth
// @Summary Gerar um token JWT de testes
// @Description Gera um token JWT de testes válido por 1 dia.
// @Produce json
// @Success 200 {object} viewmodels.Token
// @Failure 400,500 {object} viewmodels.Error
// @Router /gerar-token-jwt [get]

func (h *GerarTokenHandler) GerarTokenJWT(c *gin.Context) {

	token, err := h.GerarTokenInterface.GenerateTokenJWT()
	if err != nil {
		//util.GravarErroNoSentry(err, c)
		c.AbortWithStatusJSON(500, erroAoGerarToken)
		return
	}

	c.JSON(200, token)

}
