package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthJWT(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("No Authorization Header", func(t *testing.T) {
		router := gin.New()
		router.Use(AuthJWT("3dc3abf6a7ce3a3aa681a7873e33c9dc9f9d30159898cf623bff50d6814d8075"))
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.Contains(t, resp.Body.String(), "Authorization header não encontrado")
	})

	t.Run("Invalid Token Format", func(t *testing.T) {
		router := gin.New()
		router.Use(AuthJWT("3dc3abf6a7ce3a3aa681a7873e33c9dc9f9d30159898cf623bff50d6814d8075"))
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Add("Authorization", "Bearer")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.Contains(t, resp.Body.String(), "Token inválido")
	})

	t.Run("Valid Token", func(t *testing.T) {
		router := gin.New()
		router.Use(AuthJWT("3dc3abf6a7ce3a3aa681a7873e33c9dc9f9d30159898cf623bff50d6814d8075"))
		router.GET("/end-point-privado", func(c *gin.Context) {
			c.Status(200)
		})

		// // Assuming "valid_jwt_token" is a valid JWT for testing purposes
		// validToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlfQ.nRBsUwQPNEV8tnU_qfc5xRt5PwfcSMYDw3sUFyKBAts"
		// req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		// req.Header.Add("Authorization", validToken)
		// resp := httptest.NewRecorder()
		// router.ServeHTTP(resp, req)

		// // Print the response status code and body
		// fmt.Printf("Status code: %d\n", resp.Code)
		// fmt.Printf("Response body: %s\n", resp.Body.String())

		// assert.Equal(t, 200, nil)
	})
}
