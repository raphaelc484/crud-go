# Crud-Go

Este é um projeto CRUD simples em Go com PostgreSQL como banco de dados. O projeto utiliza Docker para facilitar a configuração e execução do banco de dados PostgreSQL.

## Requisitos

- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/install/) instalados.

## Como executar o projeto

1. Clone este repositório para a sua máquina local:

    ```bash
    git clone https://github.com/raphaelc484/crud-go
    ```

2. Entre no diretório do projeto:

    ```bash
    cd crud-go
    ```

3. Inicie o banco de dados PostgreSQL usando Docker Compose:

    ```bash
    docker-compose up -d
    ```

    Isso irá iniciar o serviço do banco de dados PostgreSQL em um contêiner Docker. O banco de dados estará disponível em `localhost:5432`.

4. Execute o projeto:

    ```bash
    go run main.go
    ```

    O servidor da API será iniciado em `http://localhost:8080`.

## Documentação

A documentação da API está disponível em `http://localhost:8080/swagger/index.html` após iniciar o servidor.

## Como parar o banco de dados

Para parar o banco de dados PostgreSQL, execute o seguinte comando:

```bash
docker-compose down