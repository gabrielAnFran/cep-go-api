package healthcheck_controller

import (
	"cep-gin-clean-arch/models"

	"github.com/gin-gonic/gin"
)

const (
	Up   = "up"
	Down = "down"
)

// @Summary      Verifica a saúde da API.
// @Description  Verifica a saúde da API. Retornando se a mesma está no ar.
// @Tags         Health Check
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.HealthCheck
// @Router       /health-check [get]
// @security 	 BasicAuth
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		res := models.HealthCheck{Status: Up}
		c.JSON(200, res)
	}
}
