# Order Creation API

Este projeto é uma API de criação de ordens construída com Go, seguindo os princípios da Clean Architecture.

## Requisitos

- Go 1.17 ou superior
- Docker (para executar o container MYSQL)

## Estrutura do Projeto

A estrutura do projeto segue os princípios da Clean Architecture, separando responsabilidades em diferentes camadas:

- **cmd**: Contém o ponto de entrada da aplicação.
- **domain**: Contém as entidades e interfaces principais.
- **usecase**: Contém os casos de uso da aplicação.
- **interface**: Contém as interfaces de entrada/saída da aplicação, como controladores HTTP e repositórios.
- **infra**: Contém as implementações concretas de interfaces, como repositórios e serviços externos.

## Instalação

1. Clone o repositório:

    ```sh
    git clone https://github.com/eduardoSantosbh/cleanarch-go.git
    cd cleanarch-go
    ```

2. Instale as dependências:

    ```sh
    go mod download
    ```

3. Subir o container do MYSQL:

    ```sh
    docker-compose up -d
    ```

## Execução

### Executando localmente

Para executar a aplicação localmente, utilize o comando:

```sh
go run cmd/main.go
