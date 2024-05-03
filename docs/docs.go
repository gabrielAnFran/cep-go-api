// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Suporte da API",
            "email": "antunes.f.gabriel@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cep/{cep}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Endpoint para buscar um CEP em um repositório",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CEP"
                ],
                "summary": "Buscar um CEP em um repositório",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CEP a ser buscado",
                        "name": "cep",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retorna o CEP encontrado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/usecase.BuscarCepOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Erro ao buscar o CEP",
                        "schema": {
                            "$ref": "#/definitions/models.CEPErrorResponse"
                        }
                    }
                }
            }
        },
        "/gerar-token": {
            "post": {
                "security": [
                    {
                        "": [
                            ""
                        ]
                    }
                ],
                "description": "Gera um token JWT para ser utilizado na requisicão de CEP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Gerar um token JWT",
                "parameters": [
                    {
                        "description": "Token Login Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TokenLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TokenLoginRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.TokenErrorResponse"
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Verifica a saúde da API. Retornando se a mesma está no ar.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Check"
                ],
                "summary": "Verifica a saúde da API.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/HealthCheck"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "HealthCheck": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "up"
                }
            }
        },
        "models.CEPErrorResponse": {
            "type": "object",
            "properties": {
                "cep": {
                    "type": "string",
                    "example": "00000000"
                },
                "error": {
                    "type": "string",
                    "example": "CEP inválido"
                }
            }
        },
        "models.TokenErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Ocorreu um erro ao gerar o token"
                }
            }
        },
        "models.TokenLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "senha": {
                    "type": "string",
                    "example": "qualquerumamenosabre123"
                }
            }
        },
        "usecase.BuscarCepOutputDTO": {
            "type": "object",
            "properties": {
                "bairro": {
                    "type": "string",
                    "example": "Inhaúma"
                },
                "cidade": {
                    "type": "string",
                    "example": "Rio de Janeiro"
                },
                "estado": {
                    "type": "string",
                    "example": "RJ"
                },
                "rua": {
                    "type": "string",
                    "example": "Rua José dos Reis"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "Documentação detalhada e exemplos sobre como usar a Especificação OpenAPI para descrever sua API.",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Desafio CEP",
	Description:      "A API Desafio CEP fornece endpoints para buscar códigos postais (CEPs) em um repositório, gerar um token JWT para autenticação e verificar a saúde da API. Esta API suporta autenticação básica, retorna respostas em formato JSON e adere à Especificação OpenAPI.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
