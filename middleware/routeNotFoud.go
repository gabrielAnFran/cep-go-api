package middlewares

import "github.com/gin-gonic/gin"

// Adicione esta nova função para lidar com rotas não encontradas
func HandleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"status":  404,
		"message": "Rota não encontrada",
	})
}
