test:
	go test -v internal/infra/database/*.go 
	go test -v middleware/*.go
	go test -v internal/usecase/*.go
	go test -v controllers/cep/*.go
	go test -v controllers/healthCheck/*.go
	go test -v controllers/token/*go
	go test -v internal/entity/*.go

# Decidi colocar um por um ao invés de ./... para ter mais controle sobre o que está sendo testado e também para mapear testes que devem ser feitos