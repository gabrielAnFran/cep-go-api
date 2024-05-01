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
		router.Use(AuthJWT())
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
		router.Use(AuthJWT())
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

	// t.Run("Valid Token", func(t *testing.T) {
	// 	router := gin.New()
	// 	router.Use(AuthJWT())
	// 	router.GET("/end-point-privado", func(c *gin.Context) {
	// 		c.Status(200)
	// 	})

	// 	// Assuming "valid_jwt_token" is a valid JWT for testing purposes
	// 	validToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ2MDA4OTcsImlhdCI6MTcxNDU5NzI5NywianRpIjoiMTc2NTlCRUQtNjMzNS00QTQxLTRENjYtRUY3OTQ4N0YwRTZFIiwibmJmIjoxNzE0NTk3Mjk3LCJzdWIiOjF9.Yx_bM5wsXbOpINZUXQOZNdBgG8QvvLwWgV4Rm1tGJG8"
	// 	req, _ := http.NewRequest(http.MethodGet, "/end-point-privado", nil)
	// 	req.Header.Add("Authorization", validToken)
	// 	resp := httptest.NewRecorder()
	// 	router.ServeHTTP(resp, req)

	// 	assert.Equal(t, 200, resp.Code)
	// })
}
