# Esta etapa no Dockerfile utiliza a imagem base do Golang para compilar o código-fonte Go.
# O diretório de trabalho é definido como /app.
# Os arquivos go.mod e go.sum são copiados para o diretório de trabalho e as dependências são baixadas.
# Todo o código-fonte é copiado para o container.
# O comando go build é executado para compilar o código-fonte Go e gerar o executável cep-api-desafio em /output.
FROM golang:1.18.2-alpine3.15 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /output/cep-api-desafio

# Esta etapa no Dockerfile utiliza a imagem base Alpine para criar a imagem final.
# Define a variável de ambiente HTTP_PORT como 8080.
# Copia o executável cep-api-desafio gerado na etapa anterior para o diretório raiz do container.
# Expõe a porta 8080.
# Define o comando a ser executado quando o container for iniciado como "/cep-api-desafio".
FROM alpine:3.15

COPY --from=builder /output/cep-api-desafio /cep-api-desafio

EXPOSE 8080

CMD ["/cep-api-desafio"]
