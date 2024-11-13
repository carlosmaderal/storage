# Usando uma imagem base do Alpine Linux, que é leve
FROM alpine:latest

# Instalando dependências necessárias para rodar o binário
RUN apk update && apk add --no-cache libc6-compat

# Montando o volume de storage
VOLUME /srv/projetos/storage

# Definindo o comando de execução
CMD ["/srv/projetos/storage/main-linux"]
