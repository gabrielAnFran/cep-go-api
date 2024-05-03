package token_controller

import (
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/models"
	"cep-gin-clean-arch/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type GerarTokenHandler struct {
	GerarTokenInterface entity.GerarTokenInterface
}

func NewGerarTokenHandler(gerarTokenRepository entity.GerarTokenInterface) *GerarTokenHandler {
	return &GerarTokenHandler{
		GerarTokenInterface: gerarTokenRepository,
	}
}

const (
	erroComDadosRequisicao = "Ocorreu um erro ao receber o corpo da requisição. Verifique se os campos foram informados corretamente."
)

// @Summary      Gerar um token JWT
// @Description  Gera um token JWT para ser utilizado na requisicão de CEP
// @Tags         Token
// @Accept       json
// @Produce      json
// @security 	 []
// @Param   req     body    models.TokenLoginRequest   true    "Token Login Request"
// @Router       /gerar-token [post]
// @Failure      500  {object}  models.TokenErrorResponse
// @Success      200 {object}   models.TokenLoginRequest
func (h *GerarTokenHandler) GerarTokenJWT(c *gin.Context) {

	var req models.TokenLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GravarErroNoSentry(err, c)
		errToken := models.TokenErrorResponse{
			Error: erroComDadosRequisicao,
		}
		c.JSON(400, errToken)
		return
	}

	if req.Email == "" || req.Senha == "" {
		errToken := models.TokenErrorResponse{
			Error: erroComDadosRequisicao,
		}
		utils.GravarErroNoSentry(errors.New(errToken.Error), c)
		c.JSON(400, errToken)
		return
	}

	token, err := h.GerarTokenInterface.GenerateTokenJWT(req)
	if err != nil {
		utils.GravarErroNoSentry(err, c)
		errToken := models.TokenErrorResponse{
			Error: "Ocorreu um erro inesperado ao gerar o token",
		}
		c.JSON(500, errToken)
		return
	}

	tokenResponse := models.TokenLoginResponse{
		Token: token,
	}
	c.JSON(200, tokenResponse)

}
