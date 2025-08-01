# 🚀 Guia de Teste dos Endpoints - Echo Playground

## 🎉 **STATUS DOS TESTES**

### ✅ **Endpoints Testados com Sucesso:**
- ✅ `GET /api/v1/` - API Principal
- ✅ `GET /api/v1/hello/Natalia` - Saudação Personalizada
- ✅ `GET /api/v1/products` - Listagem de Produtos
- ✅ `POST /api/v1/users` - Data Binding (Criar Usuário)
- ✅ `GET /api/v1/search?q=echo&category=framework&limit=5` - Query Parameters
- ✅ `GET /api/v1/xml` - Resposta XML
- ✅ `GET /api/v1/protected/profile` - Autenticação JWT (bloqueado sem token)
- ✅ `POST /api/v1/products` - Criação de Produto
- ✅ `GET /api/v1/stream` - Streaming de Dados

### 🔧 **Próximos Testes Sugeridos:**
- 🔄 `GET /api/v1/html` - Renderização HTML
- 🔄 `POST /api/v1/upload` - Upload de Arquivo
- 🔄 `GET /api/v1/download/:filename` - Download de Arquivo
- 🔄 `GET /api/v1/ws` - WebSocket (simulado)
- 🔄 `PUT /api/v1/products/:id` - Atualizar Produto
- 🔄 `DELETE /api/v1/products/:id` - Deletar Produto

## 📚 Documentação Swagger
Acesse: **http://localhost:8080/api/v1/docs**

## 🔧 Endpoints Disponíveis

### 1. **Informações Gerais**
```bash
# GET /api/v1/ - Informações sobre o framework
curl http://localhost:8080/api/v1/
```

### 2. **Saudação Personalizada**
```bash
# GET /api/v1/hello/:name - Saudação personalizada
curl http://localhost:8080/api/v1/hello/João
curl http://localhost:8080/api/v1/hello/Maria
```

### 3. **Renderização HTML**
```bash
# GET /api/v1/html - Página HTML renderizada
curl http://localhost:8080/api/v1/html
# Ou abra no navegador: http://localhost:8080/api/v1/html
```

### 4. **Resposta XML**
```bash
# GET /api/v1/xml - Resposta em formato XML
curl http://localhost:8080/api/v1/xml
```

### 5. **Data Binding - Criar Usuário**
```bash
# POST /api/v1/users - Criar usuário (data binding)
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "age": 30
  }'
```

### 5.1. **Login - Gerar Token JWT**
```bash
# POST /api/v1/login - Gerar token para autenticação
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "123456"
  }'
```

### 6. **Query Parameters - Busca**
```bash
# GET /api/v1/search - Busca com query parameters
curl "http://localhost:8080/api/v1/search?q=echo&category=framework&limit=5"
```

### 7. **Upload de Arquivo**
```bash
# POST /api/v1/upload - Upload de arquivo
echo "Conteúdo do arquivo teste" > teste.txt
curl -X POST http://localhost:8080/api/v1/upload \
  -F "file=@teste.txt"
```

### 8. **Download de Arquivo**
```bash
# GET /api/v1/download/:filename - Download de arquivo
curl http://localhost:8080/api/v1/download/exemplo.txt
```

### 9. **Streaming de Dados**
```bash
# GET /api/v1/stream - Streaming de dados
curl http://localhost:8080/api/v1/stream
```

### 10. **WebSocket (Simulado)**
```bash
# GET /api/v1/ws - WebSocket (simulado)
curl http://localhost:8080/api/v1/ws
```

## 🛡️ Endpoints Protegidos (Requer Autenticação)

### 11. **Perfil do Usuário**
```bash
# GET /api/v1/protected/profile - Perfil do usuário
# Sem token (deve retornar erro 401)
curl http://localhost:8080/api/v1/protected/profile

# Com token válido (funciona!)
curl http://localhost:8080/api/v1/protected/profile \
  -H "Authorization: Bearer valid-token"
```

**🔑 Token para Teste:** `Bearer valid-token`

## 🛍️ CRUD Completo de Produtos

### 12. **Listar Produtos**
```bash
# GET /api/v1/products - Listar todos os produtos
curl http://localhost:8080/api/v1/products
```

### 13. **Obter Produto Específico**
```bash
# GET /api/v1/products/:id - Obter produto específico
curl http://localhost:8080/api/v1/products/1
```

### 14. **Criar Produto**
```bash
# POST /api/v1/products - Criar novo produto
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Smartphone",
    "description": "Smartphone de última geração",
    "category": "Eletrônicos",
    "price": 1299.99
  }'
```

### 15. **Atualizar Produto**
```bash
# PUT /api/v1/products/:id - Atualizar produto existente
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Gaming",
    "description": "Laptop para jogos de alta performance",
    "category": "Eletrônicos",
    "price": 3999.99
  }'
```

### 16. **Deletar Produto**
```bash
# DELETE /api/v1/products/:id - Deletar produto
curl -X DELETE http://localhost:8080/api/v1/products/1
```

## 🎯 Testes com Diferentes Métodos HTTP

### Testando Métodos Não Permitidos
```bash
# Tentar GET em endpoint POST (deve retornar 405)
curl http://localhost:8080/api/v1/users

# Tentar POST em endpoint GET (deve retornar 405)
curl -X POST http://localhost:8080/api/v1/
```

### Testando Endpoints Inexistentes
```bash
# Endpoint que não existe (deve retornar 404)
curl http://localhost:8080/api/v1/inexistente
```

## 📊 Recursos Demonstrados

- ✅ **Router Otimizado**: Roteamento eficiente sem alocação dinâmica
- ✅ **Middleware Global**: Logger, Recover, CORS
- ✅ **Middleware Customizado**: Logging customizado e autenticação JWT
- ✅ **Data Binding**: Automático para JSON, XML e form-data
- ✅ **Múltiplos Formatos**: JSON, XML, HTML
- ✅ **Upload/Download**: Gerenciamento de arquivos
- ✅ **Query Parameters**: Extração e processamento
- ✅ **Path Parameters**: Parâmetros dinâmicos nas rotas
- ✅ **Tratamento de Erros**: Centralizado e customizado
- ✅ **Autenticação JWT**: Middleware de proteção
- ✅ **CRUD Completo**: Operações completas de Create, Read, Update, Delete
- ✅ **Templates HTML**: Renderização dinâmica
- ✅ **Streaming**: Transmissão de dados em tempo real
- ✅ **Arquitetura Modular**: Separação clara de responsabilidades

## 🔗 Links Úteis

- **API Principal**: http://localhost:8080/api/v1/
- **Documentação Swagger**: http://localhost:8080/api/v1/docs
- **Página HTML**: http://localhost:8080/api/v1/html
- **Framework Echo**: https://echo.labstack.com

## 🚀 Para Parar o Servidor
Pressione `Ctrl+C` no terminal onde o servidor está executando.
