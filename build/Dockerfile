# 1. Build Stage
FROM golang:1.21 AS builder

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar o go.mod e go.sum antes de baixar as dependências
COPY go.mod go.sum ./

# Baixar as dependências do projeto
RUN go mod download

# Copiar todo o código do projeto para dentro do container
COPY . .

# Compilar o binário Go
RUN go build -o /myapp ./cmd/app1

# 2. Run Stage
FROM alpine:latest

# Instalar ca-certificates para o funcionamento HTTPS
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho no container
WORKDIR /root/

# Copiar o binário compilado da etapa anterior
COPY --from=builder /myapp .

# Expor a porta usada pela aplicação
EXPOSE 8080

# Comando padrão para executar o binário
CMD ["./myapp"]