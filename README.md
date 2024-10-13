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

Essa arquitetura torna os ajustes ou novas implementações menos sofridas, onde as camadas se comunicam mas não estão "presas" umas as outras, tendo cada camada sua própria resposabilidade. Um outro fator é que possibilita testes desacoplados.

---

## Tecnologias Utilizadas

| Tecnologia | Descrição                                                                                   |
|------------|---------------------------------------------------------------------------------------------|
| **Go**        | Linguagem de programação compilada e de tipagem estática, projetada para simplicidade e eficiência. |
| **Gin**         | Um framework web escrito em Go que é usado para construir APIs de alta performance com um mínimo de recursos. |
| **Swaggo**      | Ferramenta para Go que gera automaticamente documentação de API RESTful com Swagger. |
| **Sentry**      | Serviço que ajuda a detectar, monitorar e corrigir falhas em tempo real em toda a stack da app. |
| **Mockery**     | Ferramenta para gerar mocks para testes em Go, facilitando a criação de unit tests. |
| **GitHub Actions**        | Ferramenta utilizada para rodar os testes quando um PR é submetido para a branch `main`.  |
| **Supabase**        | Alternativa de código aberto ao Firebase. possibilita a instanciação de um banco de dados Postgres, Autenticação e autorização. No contexto desse projeto, foi utilizado para a criação de uma tabela de Cep e uma função para buscar um endereço por CEP.  |

---

## Rodando o projeto

### Pré-requisitos

- Docker
- Docker Compose

### Instruções

> **Importante**: O arquivo `.env.example` mostra como deve ser o `.env` do projeto, que é necessário para rodar a aplicação. Certifique-se de criar um arquivo `.env` baseado neste exemplo antes de iniciar.

1. Clone o repositório
2. Copie o arquivo `.env.example` para `.env` e ajuste as variáveis conforme necessário
3. Execute o comando:
   ```
   docker-compose up
   ```

> **Nota**: Este comando irá construir a imagem Docker do projeto, instalar todas as dependências (incluindo Swaggo), gerar a documentação Swagger e iniciar a aplicação. Os testes podem ser executados dentro do container Docker. Uma GitHub Action também roda os testes automaticamente quando um PR é submetido para a branch **main**.

---

## Lógica da API de Busca de CEP

A API de busca de CEP implementa uma lógica flexível para encontrar informações de endereço com base em um CEP fornecido. O principal componente desta lógica está no método `Execute` da estrutura `BuscarCEPuseCase`. Aqui estão os pontos principais:

1. **Entrada flexível**: a API aceita um CEP como entrada, que pode estar incompleto ou parcialmente preenchido, desde que contenha o numero de digitos correto referente a um CEP.

2. **Busca iterativa**: o sistema realiza até 8 tentativas de busca, adicionando zeros à direita do CEP em cada iteração se a busca anterior falhar.

3. **Manipulação de erros**: 
   - Se o repositório interno retornar um erro "CEP não encontrado", o sistema tenta buscar o CEP em uma API externa.
   - Se a busca externa também falhar ou ocorrer qualquer outro tipo de erro, a execução é interrompida e o erro é retornado.


4. **Resultado da busca**: 
   - Se um CEP válido for encontrado, a função retorna um `BuscarCepOutputDTO` contendo as informações do endereço (rua, bairro, cidade e estado).
   - Se após todas as tentativas o CEP não for encontrado, a função retorna um erro indicando que o CEP não foi localizado.

5. **Abstração do repositório**: a função utiliza uma interface `CEPRepositoryInterface` para realizar a busca, permitindo flexibilidade na implementação do repositório de dados.

6. **Fallback para API externa**: se o CEP não for encontrado no repositório interno, a API tenta buscar o CEP em um serviço externo, aumentando a chance de sucesso na busca por endereços.

Esta abordagem permite que a API seja mais tolerante, aumentando a chance de encontrar um endereço correspondente mesmo quando o usuário não fornece um CEP completo ou preciso.

---

## Licença

Este projeto é distribuído sob a licença MIT License, que é uma licença de código aberto. Isso significa que qualquer pessoa pode livremente usar, modificar e distribuir o código-fonte, desde que as condições da licença MIT License sejam respeitadas. Para mais detalhes sobre o que é permitido ou não sob esta licença, visite [MIT License](https://opensource.org/license/mit).

A escolha desta licença visa promover uma colaboração aberta e o uso livre do software, permitindo que a comunidade contribua e beneficie-se das melhorias continuamente.

---

## Exemplos de uso

Exemplos de uso podem ser encontrados no projeto em `/test/postman`, onde se encontra a collection do postman e em `/test/`, onde se encontra os arquivos `.http` que podem ser usados para testar diretamente da IDE.

---

## Sentry 

### Issues 

![issues](issues.png)

### Performance

![performance](performance.png)

---

## HTTP

1. Ao digitar uma URL no browser, primeiramente, o endereço deve ser transformado em um endereço de IP. Isso acontece atravas de DNS Resolution.
2. Após a resolução de DNS, o navegador estabelece uma conexão segura (ou nao, depende) com o servidor por meio de protocolos de segurança como o HTTPS, que utiliza criptografia para proteger a comunicação entre o cliente e o servidor. O HTTPS garante a confidencialidade e integridade dos dados transmitidos, evitando que informações sensíveis sejam interceptadas por terceiros. Além disso, certificados SSL/TLS são utilizados para autenticar a identidade do servidor, garantindo que o navegador está se comunicando com o servidor correto.
3. Quando um servidor recebe uma solicitação, ele pode responder com um código de status de redirecionamento, como o código 301 (Permanent Redirect) ou 302 (Temporary Redirect), ou acessar destino diretamente. Caso haja um redirect, o navegador então segue o redirecionamento e faz uma nova solicitação para o novo endereço fornecido pelo servidor. Redirects podem ser usados para:
   1. Direcionar para uma página atualizada
   2. Corrigir URLs incorretas
   3. Redirecionar tráfego de um site para outro
4. Depois desse processo todo, a request acessa o servidor que processa a request, executa a lógica e gera a resposta apropriada, que retorna para o cliente.

```mermaid
flowchart TB
    A[Digitar uma URL no browser] --> B[Transforma o endereço em um endereço de IP - DNS Resolution]
    B --> C[Estabelece uma conexão segura com o servidor usando protocolos de segurança como HTTPS]
    C --> D[Utiliza criptografia para proteger a comunicação entre cliente e servidor]
    D --> E[Garante a confidencialidade e integridade dos dados transmitidos e utiliza certificados SSL/TLS para autenticar a identidade do servidor]
    E --> F[Recebe uma resposta do servidor com um código de status de redirecionamento, como 301 ou 302 - ou acessa o destino diretamente]
    F --> A
    F --> G[Caso redirect: Segue o redirecionamento para o novo endereço fornecido pelo servidor]
    G --> H[Realiza uma nova solicitação para o novo local, conforme indicado pelo servidor]
    H --> I[Utiliza redirecionamento para direcionar usuários para páginas atualizadas, corrigir URLs incorretas ou redirecionar tráfego entre sites]
    I --> A
```





