#!/bin/bash

# Script para testar a API usando exemplos JSON
# Uso: ./scripts/test_with_examples.sh

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configurações
BASE_URL="http://localhost:8080"
API_BASE="$BASE_URL/api/v1"
EXAMPLES_DIR="examples"

echo -e "${BLUE}🧪 Testando API com exemplos JSON${NC}"
echo "=================================="

# Função para fazer requisição e verificar resposta
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data_file=$3
    local expected_status=$4
    local description=$5

    echo -e "\n${YELLOW}📝 $description${NC}"
    echo "Endpoint: $method $endpoint"

    if [ -n "$data_file" ]; then
        echo "Data: $data_file"
        response=$(curl -s -w "\n%{http_code}" -X "$method" \
            -H "Content-Type: application/json" \
            -d "@$data_file" \
            "$API_BASE$endpoint")
    else
        response=$(curl -s -w "\n%{http_code}" -X "$method" \
            "$API_BASE$endpoint")
    fi

    # Separar resposta e status code
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)

    if [ "$http_code" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ Sucesso (Status: $http_code)${NC}"
        echo "Resposta:"
        echo "$body" | jq '.' 2>/dev/null || echo "$body"
    else
        echo -e "${RED}❌ Falha (Esperado: $expected_status, Recebido: $http_code)${NC}"
        echo "Resposta:"
        echo "$body" | jq '.' 2>/dev/null || echo "$body"
    fi
}

# Função para testar endpoint protegido
test_protected_endpoint() {
    local method=$1
    local endpoint=$2
    local description=$3

    echo -e "\n${YELLOW}🔒 $description${NC}"
    echo "Endpoint: $method $endpoint"

    # Teste sem token (deve falhar)
    echo "Testando sem token..."
    response=$(curl -s -w "\n%{http_code}" -X "$method" \
        "$API_BASE$endpoint")

    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)

    if [ "$http_code" -eq 401 ]; then
        echo -e "${GREEN}✅ Proteção funcionando (Status: $http_code)${NC}"
    else
        echo -e "${RED}❌ Proteção falhou (Esperado: 401, Recebido: $http_code)${NC}"
    fi

    # Teste com token válido
    echo "Testando com token válido..."
    response=$(curl -s -w "\n%{http_code}" -X "$method" \
        -H "Authorization: Bearer valid-token" \
        "$API_BASE$endpoint")

    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)

    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✅ Autenticação funcionando (Status: $http_code)${NC}"
        echo "Resposta:"
        echo "$body" | jq '.' 2>/dev/null || echo "$body"
    else
        echo -e "${RED}❌ Autenticação falhou (Esperado: 200, Recebido: $http_code)${NC}"
        echo "Resposta:"
        echo "$body" | jq '.' 2>/dev/null || echo "$body"
    fi
}

# Verificar se o servidor está rodando
echo -e "${BLUE}🔍 Verificando se o servidor está rodando...${NC}"
if ! curl -s "$BASE_URL/api/v1/" > /dev/null; then
    echo -e "${RED}❌ Servidor não está rodando em $BASE_URL${NC}"
    echo "Execute: go run cmd/echo-playground/main.go ou ./bin/echo"
    exit 1
fi
echo -e "${GREEN}✅ Servidor está rodando${NC}"

# Testes básicos
echo -e "\n${BLUE}📋 Executando testes básicos...${NC}"
test_endpoint "GET" "/" "" 200 "Informações do framework"
test_endpoint "GET" "/hello/João" "" 200 "Saudação personalizada"
test_endpoint "GET" "/html" "" 200 "Página HTML"
test_endpoint "GET" "/xml" "" 200 "Resposta XML"
test_endpoint "GET" "/search?q=laptop&limit=5" "" 200 "Busca com query parameters"

# Testes de usuários
echo -e "\n${BLUE}👥 Executando testes de usuários...${NC}"
test_endpoint "POST" "/users" "$EXAMPLES_DIR/create_user.json" 201 "Criar usuário"

# Testes de produtos
echo -e "\n${BLUE}📦 Executando testes de produtos...${NC}"
test_endpoint "GET" "/products" "" 200 "Listar produtos"
test_endpoint "POST" "/products" "$EXAMPLES_DIR/create_product.json" 201 "Criar produto"
test_endpoint "GET" "/products/1" "" 200 "Obter produto específico"
test_endpoint "PUT" "/products/1" "$EXAMPLES_DIR/update_product.json" 200 "Atualizar produto"
test_endpoint "DELETE" "/products/1" "" 200 "Deletar produto"

# Testes de upload/download
echo -e "\n${BLUE}📁 Executando testes de arquivos...${NC}"
echo "Testando upload de arquivo..."
if [ -f "test/teste.txt" ]; then
    response=$(curl -s -w "\n%{http_code}" -X POST \
        -F "file=@test/teste.txt" \
        "$API_BASE/upload")

    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)

    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✅ Upload funcionando (Status: $http_code)${NC}"
        echo "Resposta:"
        echo "$body" | jq '.' 2>/dev/null || echo "$body"
    else
        echo -e "${RED}❌ Upload falhou (Status: $http_code)${NC}"
    fi
else
    echo -e "${YELLOW}⚠️  Arquivo test/teste.txt não encontrado, pulando teste de upload${NC}"
fi

# Testes de autenticação
echo -e "\n${BLUE}🔐 Executando testes de autenticação...${NC}"
test_protected_endpoint "GET" "/protected/profile" "Endpoint protegido"

# Testes de streaming
echo -e "\n${BLUE}🌊 Executando testes de streaming...${NC}"
echo "Testando streaming..."
response=$(curl -s -w "\n%{http_code}" "$API_BASE/stream")
http_code=$(echo "$response" | tail -n1)
if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✅ Streaming funcionando (Status: $http_code)${NC}"
else
    echo -e "${RED}❌ Streaming falhou (Status: $http_code)${NC}"
fi

# Testes de WebSocket (simulado)
echo -e "\n${BLUE}🔌 Executando testes de WebSocket...${NC}"
test_endpoint "GET" "/ws" "" 200 "WebSocket (simulado)"

# Resumo final
echo -e "\n${BLUE}📊 Resumo dos testes${NC}"
echo "=================="
echo -e "${GREEN}✅ Todos os testes concluídos!${NC}"
echo ""
echo "Para testar manualmente com exemplos JSON:"
echo "  curl -X POST -H 'Content-Type: application/json' \\"
echo "    -d @examples/create_user.json \\"
echo "    http://localhost:8080/api/v1/users"
echo ""
echo "Para ver a documentação Swagger:"
echo "  swagger serve api/swagger.yaml"
echo ""
echo "Para executar testes automatizados:"
echo "  ./scripts/test_api.sh"
