#!/bin/bash

# Script de build para o Echo Playground
# Este script compila e testa a aplicação

set -e

echo "🔨 Build do Echo Playground"
echo "=========================="

# Verificar se o Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go não está instalado"
    exit 1
fi

# Limpar builds anteriores
echo "🧹 Limpando builds anteriores..."
go clean

# Verificar dependências
echo "📦 Verificando dependências..."
go mod tidy

# Executar testes
echo "🧪 Executando testes..."
go test ./...

# Compilar aplicação
echo "🔨 Compilando aplicação..."
go build -o bin/echo ./cmd/echo-playground

echo "✅ Build concluído com sucesso!"
echo "📁 Executável criado em: bin/echo"
echo ""
echo "Para executar:"
echo "  ./bin/echo"
