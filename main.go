package main

import (
	cep_controller "cep-gin-clean-arch/controllers/cep"
	healthcheck_controller "cep-gin-clean-arch/controllers/healthCheck"
	token_controller "cep-gin-clean-arch/controllers/token"
	_ "cep-gin-clean-arch/docs"
	"cep-gin-clean-arch/internal/entity"
	"cep-gin-clean-arch/internal/infra/database"
	"cep-gin-clean-arch/internal/usecase"
	middlewares "cep-gin-clean-arch/middleware"
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @version         1.0
// @description     A API Desafio CEP fornece endpoints para buscar códigos postais (CEPs) em um repositório, gerar um token JWT para autenticação e verificar a saúde da API. Esta API suporta autenticação básica, retorna respostas em formato JSON e adere à Especificação OpenAPI.
// @termsOfService  http://swagger.io/terms/
// @title           API Desafio CEP
// @contact.name    Suporte da API
// @contact.email   antunes.f.gabriel@gmail.com
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apiKey.description Use sua chave de API para acessar esta API.
// @externalDocs.description  Documentação detalhada e exemplos sobre como usar a Especificação OpenAPI para descrever sua API.
// @externalDocs.url          https://swagger.io/resources/open-api/
// @tag             Operações de CEP
func main() {
	// Carrega variáveis de ambiente
	godotenv.Load()
	sentryInit()

	// Inicializa o gin router
	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{}))
	router.Use(cors.Default())

	// Inicializa as rotas da aplicação
	rotas(router)

	// Inicializa o servidor
	port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	fmt.Println("Servidor inicializado na porta:", os.Getenv("HTTP_PORT"))
	router.Run(port)
}

func rotas(router *gin.Engine) {
	// Rotas que necessitam de autenticação
	api := router.Group("/")
	api.Use(middlewares.AuthJWT())
	{
		// Inicializa o repositório e o handler do CEP
		repository := entity.CEPRepositoryInterface(database.NewCEPRepository())
		webCEPHandler := cep_controller.CEPWebHandler{CEPRepository: repository}
		api.GET("/cep/:cep", webCEPHandler.BuscarCEP)
		api.GET("/health-check", healthcheck_controller.HealthCheck())
	}

	// rotas que nao necessitam de autenticação
	auth := router.Group("/")
	{
		// Chamada do serviço de auth para gerar um token JWT
		jwtService := usecase.UsecaseAuth{}
		gerarTokenHandler := token_controller.GerarTokenHandler{GerarTokenInterface: jwtService}
		auth.POST("/gerar-token", gerarTokenHandler.GerarTokenJWT)
	}

	// Rota do swagger sem auth
	swagger := router.Group("/docs")
	{
		swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}

func sentryInit() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0, // Capture 100% of transactions for monitoring
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}
