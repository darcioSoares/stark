# Project API finance

Este é a estrutura da API Stark Bank (teste)

```plaintext
app/
├── cmd/                      # Entrypoints da aplicação
|   |── server/               # Pasta aonde esta o arquivo principal
│       └── main.go           # Arquivo principal do aplicativo
├── internal/                 # Código interno (não exportado para outros projetos)
│   ├── handlers/             # Handlers para rotas (controllers)
│   ├── routes/               # Definições das rotas
│   ├── services/             # Lógica de negócios
│   ├── models/               # Definição das structs (modelos de dados)
│   └── utils/                # Funções utilitárias
├── go.mod                    # Gerenciamento de dependências
├── go.sum                    # Checksum das dependências
├── README.md                 # Documentação do projeto
```
--------------------------------------------------------------------------------

# Passo a passo para rodar a aplicação

1. Clone o repositório:

git clone https://github.com/darcioSoares/stark
cd stark

### Configuração do arquivo .env

Renomeie o arquivo .env-exemplo para .env Que se encontra dentro da pasta APP

mv .env-exemplo .env

Adicione os valores necessários:

PRIVATE_KEY: Chave privada da aplicação. (Tomar cuidado com os espaços seguir exemplo do arq .env-exemplo)

ID_PROJECT: ID do projeto.

## 2. Suba os containers com Docker: (Executar dentro da Raiz do projeto)

Dentro da pasta do projeto, use o comando:

docker-compose up ou docker compose up dependendo da versão do compose

Este comando irá subir os containers necessários para a aplicação.

## 3. Rodar a aplicação sem Docker:

Caso prefira rodar sem Docker, utilize o seguinte comando:

go mod tidy (Baixar as dependencias )

go run cmd/server/main.go

-------------------------------------------------------------------------------

Observação usei ngrok http 8080 para receber webhooks

--------------------------------------------------------------------------------

## Sobre as Branches

Branch master (síncrona):

O projeto na branch master funciona de forma síncrona, processando as requisições imediatamente.

Branch rabbitmq (assíncrona):

Para rodar de forma assíncrona, utilize a branch rabbitmq.

Nesta branch, a aplicação se conecta ao RabbitMQ e usa filas para armazenar os retornos do webhook.

Como baixar e trocar para a branch rabbitmq:

git fetch origin rabbitmq ( para baixar ) git checkout -b rabbitmq origin/rabbitmq ( Isso cria e troca para a branch local rabbitmq)

Todo o restante já está configurado no arquivo.

-------------------------------------------------------

## Conclusão

Seguindo esses passos, a aplicação estará pronta para rodar, seja em modo síncrono ou assíncrono, dependendo da branch utilizada. Caso tenha dúvidas, revise o README.md ou entre em contato.

