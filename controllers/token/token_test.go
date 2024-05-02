package token_controller

import (
	"cep-gin-clean-arch/internal/usecase"
	"cep-gin-clean-arch/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBuscarTokenSucesso(t *testing.T) {
	serviceToken := new(mocks.GerarTokenInterface)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	// Simulate a POST request with request body {"email": "aa", "senha": "senha"}
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email": "aa@bb.com", "senha": "senha"}`))

	jwtService := usecase.UsecaseAuth{}
	tokenHandler := GerarTokenHandler{GerarTokenInterface: jwtService}
	serviceToken.On("GenerateTokenJWT").Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlfQ.nRBsUwQPNEV8tnU_qfc5xRt5PwfcSMYDw3sUFyKBAts", nil)
	tokenHandler.GerarTokenJWT(c)

	responseData, err := io.ReadAll(response.Body)
	assert.Equal(t, nil, err)

	assert.Equal(t, 200, response.Code)
	assert.NotEmpty(t, responseData)
}

func TestNewGerarTokenHandler(t *testing.T) {
	mockRepository := new(usecase.UsecaseAuth)
	handler := NewGerarTokenHandler(mockRepository)

	assert.NotNil(t, handler)
	assert.Equal(t, mockRepository, handler.GerarTokenInterface)
}
