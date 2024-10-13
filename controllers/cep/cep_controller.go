package cep_controller

import (
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/internal/usecase"
	"cep-gin-clean-arch/models"
	"cep-gin-clean-arch/utils"

	"github.com/gin-gonic/gin"
)

type CEPWebHandler struct {
	CEPRepository   entity.CEPRepositoryInterface
	BuscaCepExterno entity.CEPServiceInterface
}

func NewBuscarCEPHandler(buscarCEPRepository entity.CEPRepositoryInterface, buscaCepExterno entity.CEPServiceInterface) *CEPWebHandler {
	return &CEPWebHandler{
		CEPRepository:   buscarCEPRepository,
		BuscaCepExterno: buscaCepExterno,
	}
}

// @Summary      Buscar um CEP em um repositório
// @Description  Endpoint para buscar um CEP em um repositório
// @Tags         CEP
// @Accept       json
// @Produce      json
// @Param        cep   path      string  true  "CEP a ser buscado sem hífen"
// @Success      200  {object}  usecase.BuscarCepOutputDTO  "Retorna o CEP encontrado com sucesso"
// @Success      400  {object}  models.CEPErrorResponse       "Erro ao buscar o CEP"
// @Router       /cep/{cep} [get]
// @securityDefinitions.apiKey OAuth2
// @in header
// @name Authorization
func (h *CEPWebHandler) BuscarCEP(c *gin.Context) {

	cepParam := c.Param("cep")
	cep := entity.NewCep(cepParam)
	err := cep.IsValidCep(cep.Cep)
	if err != nil {
		utils.GravarErroNoSentry(err, c)
		errorResponse := models.CEPErrorResponse{
			Error:        err.Error(),
			CepInformado: cep.Cep,
		}
		c.JSON(400, errorResponse)
		return
	}

	cepBuscar := usecase.NewBuscarCEPUseCase(h.CEPRepository, h.BuscaCepExterno)
	res, err := cepBuscar.Execute(&cepParam)
	if err != nil {
		utils.GravarErroNoSentry(err, c)
		errorResponse := models.CEPErrorResponse{
			Error:        err.Error(),
			CepInformado: cep.Cep,
		}
		c.JSON(400, errorResponse)
		return
	}

	c.JSON(200, res)
}
