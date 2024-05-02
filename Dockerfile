
FROM golang:1.18.2-alpine3.15 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /output/cep-api-desafio

FROM alpine:3.15

ENV HTTP_PORT=8080

COPY --from=builder /output/cep-api-desafio /cep-api-desafio

EXPOSE 8080

# Define o comando padrão para executar o aplicativo
CMD ["/cep-api-desafio"]
