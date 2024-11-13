# Use uma imagem base com Go
FROM golang:1.23-alpine

# Defina o diretório de trabalho no container
WORKDIR /app

# Copiar o código-fonte do repositório para o container
COPY . .

# Certificar-se de que o script de build tenha permissões de execução
RUN chmod +x build_and_run.sh

# Instalar dependências Go (se necessário)
RUN go mod tidy

# Comando que será executado a cada reinício do container
CMD ["./build_and_run.sh"]
