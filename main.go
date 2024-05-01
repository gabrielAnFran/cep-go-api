package main

import (
	"cep-gin-clean-arch/configs"
	"cep-gin-clean-arch/controllers"
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/internal/infra/database"
	"cep-gin-clean-arch/internal/usecase"
	middlewares "cep-gin-clean-arch/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Initialize Gin router
	router := gin.Default()

	api := router.Group("/api")
	api.Use(middlewares.AuthJWT(configs.JWTSecret))

	jwtService := usecase.ServiceAuth{}
	gerarTokenHandler := controllers.GerarTokenHandler{GerarTokenInterface: jwtService}
	router.GET("/gerar-token-jwt", gerarTokenHandler.GerarTokenJWT)

	// Setup route

	// Create repository and handler
	repository := entity.CEPRepositoryInterface(database.NewCEPRepository())
	webCEPHandler := controllers.CEPWebHandler{CEPRepository: repository}

	router.GET("/cep/:cep", webCEPHandler.BuscarCEP)

	// Start server
	port := fmt.Sprintf(":%s", configs.HTTPPort)
	fmt.Println("Servidor inicializado na porta:", configs.HTTPPort)
	router.Run(port)

}
