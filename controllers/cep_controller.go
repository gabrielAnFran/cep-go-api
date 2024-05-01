package controllers

import (
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/internal/usecase"
	"errors"

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
		c.AbortWithStatusJSON(400, errors.New("erro"))
		return
	}

	cebBuscar := usecase.NewBuscarCEPUseCase(h.CEPRepository)
	res, err := cebBuscar.Execute(&cepParam)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, res)
}
