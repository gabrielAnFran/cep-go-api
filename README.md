<h1 align="center">
  CEP API 
</h1>

## Descrição

Este projeto visa desenvolver uma API robusta para a consulta de endereços utilizando como entrada o Código de Endereçamento Postal (CEP) do usuário. A API deve ser capaz de retornar informações detalhadas como rua, bairro, cidade e estado, além de tratar situações onde o CEP pode estar incompleto ou ser inválido.

---

## Arquitetura

![clean arch](clean-arch.jpg)

Esse projeto é uma implementaçao de clean architecture, tentando ao máximo seguir as boas práticas de desenvolvimento de APIs. 

A arquitura usada levou em consideração os seguintes componentes (ou camadas):

| Componente   | Descrição                                                                                           |
|--------------|-----------------------------------------------------------------------------------------------------|
| **Entities**    | Representa as entidades de domínio do projeto, como objetos que refletem conceitos do negócio.     |
| **Usecases**  | Contém a lógica de negócio do sistema, implementando os casos de uso específicos da aplicação.     |
| **Middleware**  | Responsável por interceptar e processar requisições antes de chegarem aos controllers.             |
| **Controllers**  | Recebem requisições, interagem com os usecases e retornam respostas adequadas para o cliente.      |
| **Utils**        | Módulo contendo funções utilitárias reutilizáveis em diferentes partes do projeto.                 |
| **Mocks**        | Utilizado para simular comportamentos de componentes externos durante testes ou desenvolvimento.  |

### Por que essa arquitetura?

Essa arquitetura torna os ajustes ou novas implementações menos sofridas, onde as camadas se comunicam mas não estão "presas" umas as outras, tendo cada camada sua própria resposabilidade.

---

## Tecnologias Utilizadas

| Tecnologia | Descrição                                                                                   |
|------------|---------------------------------------------------------------------------------------------|
| **Go**        | Linguagem de programação compilada e de tipagem estática, projetada para simplicidade e eficiência. |
| **Gin**         | Um framework web escrito em Go que é usado para construir APIs de alta performance com um mínimo de recursos. |
| **Swaggo**      | Ferramenta para Go que gera automaticamente documentação de API RESTful com Swagger. |
| **Sentry**      | Serviço que ajuda a detectar, monitorar e corrigir falhas em tempo real em toda a stack da app. |
| **Mockery**     | Ferramenta para gerar mocks para testes em Go, facilitando a criação de unit tests. |
| **GitHub Actions**        | Ferramenta utilizada para rodar os testes quando um PR é submetido para a brandh `main`.  |
---

## Licença

Este projeto é distribuído sob a licença MIT License, que é uma licença de código aberto. Isso significa que qualquer pessoa pode livremente usar, modificar e distribuir o código-fonte, desde que as condições da licença MIT License sejam respeitadas. Para mais detalhes sobre o que é permitido ou não sob esta licença, consulte o arquivo LICENSE incluído neste repositório ou visite [MIT License](https://opensource.org/license/mit).

A escolha desta licença visa promover uma colaboração aberta e o uso livre do software, permitindo que a comunidade contribua e beneficie-se das melhorias continuamente.