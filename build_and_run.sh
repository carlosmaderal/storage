#!/bin/bash

# Definindo o nome do binário
BINARY_NAME="main"

# Compilando o projeto Go para Linux
echo "Compilando o projeto Go..."
go build -o main .

# Verificando se a compilação foi bem-sucedida
if [ $? -eq 0 ]; then
  echo "Compilação concluída com sucesso!"
else
  echo "Erro na compilação!"
  exit 1
fi

# Executando o binário gerado
echo "Executando o binário..."
./main
