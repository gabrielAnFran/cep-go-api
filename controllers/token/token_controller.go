package token_controller

import (
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/models"
	"cep-gin-clean-arch/utils"

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

// @Summary      Gerar um token JWT
// @Description  Gera um token JWT para ser utilizado na requisicão de CEP
// @Tags         Token
// @Accept       json
// @Produce      json
// @security 	 BasicAuth
// @Success      200  {object}  string
// @Router       /gerar-token [post]
// @Success      500  {object}  models.TokenErrorResponse
func (h *GerarTokenHandler) GerarTokenJWT(c *gin.Context) {

	var req models.TokenLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GravarErroNoSentry(err, c)
		c.JSON(400, gin.H{
			"error": "Ocorreu um erro ao receber dados da requisição. Verifique se os campos estão corretos.",
		})
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

	c.JSON(200, gin.H{
		"token": token,
	})

}
