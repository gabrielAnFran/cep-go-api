package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleNotFound(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.NoRoute(HandleNotFound)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/cepppp", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(t, float64(404), response["status"])
	assert.Equal(t, "Rota não encontrada", response["message"])
}
