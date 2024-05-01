package main

import (
	cep_controller "cep-gin-clean-arch/controllers/cep"
	token_controller "cep-gin-clean-arch/controllers/token"
	"cep-gin-clean-arch/docs"
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

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Carrega variáveis de ambiente
	godotenv.Load()

	// Inicializa o Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0, // Capture 100% of transactions for monitoring
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Inicializa o gin router
	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{})) // Anexa o middleware do Sentry ao gin
	router.Use(cors.Default())                     // Viabiliza requisições de diferentes origens

	// Rotas que necessitam de autenticação
	api := router.Group("/")
	api.Use(middlewares.AuthJWT())

	// rotas que nao necessitam de autenticação
	auth := router.Group("/")

	// Inicializa o serviço de autenticação e o handler
	jwtService := usecase.ServiceAuth{}
	gerarTokenHandler := token_controller.GerarTokenHandler{GerarTokenInterface: jwtService}
	auth.POST("/gerar-token", gerarTokenHandler.GerarTokenJWT)

	// Inicializa o repositório de CEP e o handler
	repository := entity.CEPRepositoryInterface(database.NewCEPRepository())
	webCEPHandler := cep_controller.CEPWebHandler{CEPRepository: repository}
	api.GET("/cep/:cep", webCEPHandler.BuscarCEP)

	// Configuração do Swagger
	docs.SwaggerInfo.Title = "API de Consulta de CEP"
	docs.SwaggerInfo.Description = "API de consulta de CEP utilizando Clean Architecture e Golang"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Serve Swagger UI
	router.GET("/swagger/", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicializa o servidor
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println("Servidor inicializado na porta:", os.Getenv("PORT"))
	router.Run(port)
}
