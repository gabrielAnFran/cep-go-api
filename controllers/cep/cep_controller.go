package controllers

import (
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/internal/usecase"

	"github.com/gin-gonic/gin"
)

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
		c.AbortWithStatusJSON(400, gin.H{
			"error":         err.Error(),
			"Cep informado": cep.Cep})
		return
	}

	cepBuscar := usecase.NewBuscarCEPUseCase(h.CEPRepository)
	res, err := cepBuscar.Execute(&cepParam)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, res)
}
