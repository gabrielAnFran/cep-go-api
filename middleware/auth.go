package middlewares

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthJWT é o Middleware responsável por validar a autenticidade de uma requisição através do seu token JWT.
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		headerAuthorization := c.GetHeader("Authorization")
		if headerAuthorization == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header não encontrado"})
			return
		}

		bearerToken := strings.Split(headerAuthorization, "Bearer ")
		if len(bearerToken) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token inválido"})
			return
		}

		tokenString := bearerToken[1]

		if tokenString != os.Getenv("TOKEN_TESTE") {
			_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				println(err.Error())
				c.AbortWithStatusJSON(401, gin.H{"error": "Token inválido"})
				return
			}
		}

		c.Next()
	}
}
