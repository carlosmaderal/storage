# Use uma imagem base com Go
FROM golang:1.23-alpine

# Defina o diretório de trabalho no container
WORKDIR /app

# Instalar dependências Go (se necessário)
RUN go mod tidy

# Comando que será executado a cada reinício do container
CMD ["./build_and_run.sh"]
