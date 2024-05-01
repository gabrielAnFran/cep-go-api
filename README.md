<h1 align="center">
  CEP API 
</h1>

## Descrição
Breve descrição do que o projeto faz e qual problema ele visa resolver.//TODO

## Motivação
Explique por que este projeto foi iniciado, qual necessidade ele atende.//TODO

## Arquitetura

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
Justifique as escolhas arquiteturais, incluindo os benefícios de usar tal configuração. //TODO

## Tecnologias Utilizadas
Liste as tecnologias principais utilizadas no projeto e uma breve descrição de por que foram escolhidas. //TODO

### Swagger
Explique como o Swagger é utilizado para documentação da API e como isso beneficia o projeto. //TODO

### Sentry
Descreva como o Sentry é integrado e utilizado para monitoramento de erros e performance. //TODO

## Como Executar
Instruções passo a passo de como configurar e executar o projeto localmente. //TODO

## Contribuições
Instruções para desenvolvedores que desejam contribuir para o projeto. //TODO

## Licença

Este projeto é distribuído sob a licença MIT License, que é uma licença de código aberto. Isso significa que qualquer pessoa pode livremente usar, modificar e distribuir o código-fonte, desde que as condições da licença MIT License sejam respeitadas. Para mais detalhes sobre o que é permitido ou não sob esta licença, consulte o arquivo LICENSE incluído neste repositório ou visite [MIT License](https://opensource.org/license/mit).

A escolha desta licença visa promover uma colaboração aberta e o uso livre do software, permitindo que a comunidade contribua e beneficie-se das melhorias continuamente.