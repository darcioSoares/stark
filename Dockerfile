FROM golang:1.23

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app/ ./

# Produção
RUN if [ "$MODE" = "prod" ]; then go build -o main ./cmd/server; fi

# Configurar o comando padrão com base no modo
CMD if [ "$MODE" = "prod" ]; then ./main; else go run cmd/server/main.go; fi

EXPOSE 8080

