package healthcheck_controlle

import (
	"cep-gin-clean-arch/models"

	"github.com/gin-gonic/gin"
)

const (
	Up   = "up"
	Down = "down"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		res := models.HealthCheck{Status: Up}
		c.JSON(200, res)
	}
}
