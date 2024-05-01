package cep_controller

import (
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/internal/usecase"
	"cep-gin-clean-arch/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary      Busca um CEP
// @Description  Busca um CEP informado como parâmetro em um banco de dados mackado
// @Tags         cep
// @Accept       json
// @Produce      json
// @Param        cep   path      string  true  "CEP"
// @Success      200  {object}  usecase.BuscarCepOutputDTO
// @Router       /cep/{cep} [get]
type CEPWebHandler struct {
	CEPRepository entity.CEPRepositoryInterface
}

func NewBuscarCEPHandler(buscarCEPRepository entity.CEPRepositoryInterface) *CEPWebHandler {
	return &CEPWebHandler{
		CEPRepository: buscarCEPRepository,
	}
}

func (h *CEPWebHandler) BuscarCEP(c *gin.Context) {

	cepParam := c.Param("cep")
	cep := entity.NewCep(cepParam)
	err := cep.IsValidCep(cep.Cep)
	if err != nil {
		utils.GravarErroNoSentry(err, c)
		c.AbortWithStatusJSON(400, gin.H{
			"error":         err.Error(),
			"Cep informado": cep.Cep})
		return
	}

	cepBuscar := usecase.NewBuscarCEPUseCase(h.CEPRepository)
	res, err := cepBuscar.Execute(&cepParam)
	if err != nil {
		fmt.Println(err.Error())
		utils.GravarErroNoSentry(err, c)
		c.AbortWithStatusJSON(400, gin.H{
			"error":         err.Error(),
			"Cep informado": cep.Cep})
		return
	}

	c.JSON(200, res)
}
