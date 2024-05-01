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

	api := router.Group("/")
	api.Use(middlewares.AuthJWT(configs.JWTSecret))

	auth := router.Group("/")

	jwtService := usecase.ServiceAuth{}
	gerarTokenHandler := controllers.GerarTokenHandler{GerarTokenInterface: jwtService}
	auth.POST("/gerar-token", gerarTokenHandler.GerarTokenJWT)

	repository := entity.CEPRepositoryInterface(database.NewCEPRepository())
	webCEPHandler := controllers.CEPWebHandler{CEPRepository: repository}
	api.GET("/cep/:cep", webCEPHandler.BuscarCEP)

	port := fmt.Sprintf(":%s", configs.HTTPPort)
	fmt.Println("Servidor inicializado na porta:", configs.HTTPPort)
	router.Run(port)

}
