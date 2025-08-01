#!/bin/bash

# 🚀 Echo Playground - Script de Teste da API
# Este script demonstra todas as funcionalidades do Echo framework

echo "🚀 Echo Playground - Testando todas as funcionalidades"
echo "=================================================="
echo ""

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

BASE_URL="http://localhost:8080/api/v1"

# Função para testar endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local headers=$4

    echo -e "${BLUE}Testando: ${method} ${endpoint}${NC}"

    if [ -n "$data" ]; then
        if [ -n "$headers" ]; then
            response=$(curl -s -X ${method} -H "${headers}" -d "${data}" "${BASE_URL}${endpoint}")
        else
            response=$(curl -s -X ${method} -d "${data}" "${BASE_URL}${endpoint}")
        fi
    else
        if [ -n "$headers" ]; then
            response=$(curl -s -X ${method} -H "${headers}" "${BASE_URL}${endpoint}")
        else
            response=$(curl -s -X ${method} "${BASE_URL}${endpoint}")
        fi
    fi

    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Sucesso${NC}"
        echo "$response" | jq '.' 2>/dev/null || echo "$response"
    else
        echo -e "${RED}❌ Erro${NC}"
    fi
    echo ""
}

# Verificar se o servidor está rodando
echo -e "${YELLOW}Verificando se o servidor está rodando...${NC}"
if curl -s "${BASE_URL}/" > /dev/null; then
    echo -e "${GREEN}✅ Servidor está rodando!${NC}"
    echo ""
else
    echo -e "${RED}❌ Servidor não está rodando. Execute 'go run main.go' primeiro.${NC}"
    exit 1
fi

# 1. Informações gerais
echo -e "${YELLOW}1. Informações Gerais${NC}"
test_endpoint "GET" "/"

# 2. Demonstrações de resposta
echo -e "${YELLOW}2. Demonstrações de Resposta${NC}"
test_endpoint "GET" "/hello/João"
test_endpoint "GET" "/xml"
test_endpoint "GET" "/html"

# 3. Data Binding
echo -e "${YELLOW}3. Data Binding${NC}"
test_endpoint "POST" "/users" '{"name":"Maria Silva","email":"maria@exemplo.com","age":25}' "Content-Type: application/json"

# 4. Query Parameters
echo -e "${YELLOW}4. Query Parameters${NC}"
test_endpoint "GET" "/search?q=laptop&limit=5"

# 5. Upload de arquivo
echo -e "${YELLOW}5. Upload de Arquivo${NC}"
if [ -f "teste.txt" ]; then
    echo -e "${BLUE}Testando: POST /upload${NC}"
    response=$(curl -s -X POST -F "file=@teste.txt" "${BASE_URL}/upload")
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Sucesso${NC}"
        echo "$response" | jq '.' 2>/dev/null || echo "$response"
    else
        echo -e "${RED}❌ Erro${NC}"
    fi
    echo ""
else
    echo -e "${YELLOW}⚠️  Arquivo teste.txt não encontrado, pulando teste de upload${NC}"
    echo ""
fi

# 6. Autenticação JWT
echo -e "${YELLOW}6. Autenticação JWT${NC}"
test_endpoint "GET" "/protected/profile" "" "Authorization: Bearer valid-token"

# 7. CRUD Completo - Produtos
echo -e "${YELLOW}7. CRUD Completo - Produtos${NC}"

# Listar produtos
test_endpoint "GET" "/products"

# Obter produto específico
test_endpoint "GET" "/products/1"

# Criar produto
test_endpoint "POST" "/products" '{"name":"Teclado Mecânico","price":299.99,"description":"Teclado mecânico RGB","category":"Acessórios"}' "Content-Type: application/json"

# Atualizar produto
test_endpoint "PUT" "/products/1" '{"name":"Laptop Atualizado","price":3499.99,"description":"Laptop de alta performance atualizado","category":"Eletrônicos"}' "Content-Type: application/json"

# Deletar produto
test_endpoint "DELETE" "/products/1"

# 8. Streaming
echo -e "${YELLOW}8. Streaming${NC}"
echo -e "${BLUE}Testando: GET /stream (primeiros 3 chunks)${NC}"
curl -s -N "${BASE_URL}/stream" | head -3
echo ""
echo ""

# 9. WebSocket (simulado)
echo -e "${YELLOW}9. WebSocket (Simulado)${NC}"
test_endpoint "GET" "/ws"

# 10. Teste de erro
echo -e "${YELLOW}10. Teste de Tratamento de Erro${NC}"
echo -e "${BLUE}Testando: GET /endpoint-inexistente${NC}"
response=$(curl -s -w "%{http_code}" "${BASE_URL}/endpoint-inexistente")
http_code="${response: -3}"
if [ "$http_code" = "404" ]; then
    echo -e "${GREEN}✅ Erro tratado corretamente (404)${NC}"
else
    echo -e "${RED}❌ Erro inesperado${NC}"
fi
echo ""

echo -e "${GREEN}🎉 Todos os testes concluídos!${NC}"
echo ""
echo -e "${BLUE}📖 Para mais informações, consulte:${NC}"
echo "   - README.md"
echo "   - API_DOCUMENTATION.md"
echo ""
echo -e "${BLUE}🌐 Acesse: http://localhost:8080/api/v1/html para ver a página HTML${NC}"
