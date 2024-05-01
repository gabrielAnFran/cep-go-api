package token_controller

import (
	"cep-gin-clean-arch/mocks"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBuscarTokenSucesso(t *testing.T) {
	serviceToken := new(mocks.GerarTokenInterface)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	tokenHandler := GerarTokenHandler{GerarTokenInterface: serviceToken}
	serviceToken.On("GenerateTokenJWT").Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlfQ.nRBsUwQPNEV8tnU_qfc5xRt5PwfcSMYDw3sUFyKBAts", nil)
	tokenHandler.GerarTokenJWT(c)

	responseData, err := io.ReadAll(response.Body)
	assert.Equal(t, nil, err)

	assert.Equal(t, 200, response.Code)
	assert.NotEmpty(t, responseData)
}
