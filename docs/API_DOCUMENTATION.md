# 🚀 Echo Playground - Documentação da API

Este é um playground completo que demonstra todas as funcionalidades do Echo framework, um framework Go de alta performance, extensível e minimalista.

## 📋 Índice

- [Instalação e Execução](#instalação-e-execução)
- [Endpoints da API](#endpoints-da-api)
- [Funcionalidades Demonstradas](#funcionalidades-demonstradas)
- [Exemplos de Uso](#exemplos-de-uso)

## 🛠️ Instalação e Execução

### Pré-requisitos
- Go 1.21 ou superior
- Git

### Executar o projeto
```bash
# Clonar o repositório
git clone <repository-url>
cd echo-playground

# Instalar dependências
go mod tidy

# Executar o servidor
go run main.go
```

O servidor estará disponível em: `http://localhost:8080`

## 📚 Endpoints da API

### Base URL
```
http://localhost:8080/api/v1
```

### 1. Informações Gerais

#### GET `/`
Retorna informações sobre o framework e suas funcionalidades.

**Resposta:**
```json
{
  "success": true,
  "message": "Bem-vindo ao Echo Playground!",
  "data": {
    "framework": "Echo",
    "version": "v4",
    "features": [
      "High Performance Router",
      "Scalable REST APIs",
      "Automatic TLS",
      "HTTP/2 Support",
      "Middleware System",
      "Data Binding",
      "Data Rendering",
      "Template Support",
      "Extensible Architecture"
    ]
  }
}
```

### 2. Demonstrações de Resposta

#### GET `/hello/:name`
Retorna uma saudação personalizada.

**Parâmetros:**
- `name` (path): Nome da pessoa

**Exemplo:**
```bash
curl http://localhost:8080/api/v1/hello/João
```

**Resposta:**
```
Olá, João! Bem-vindo ao Echo!
```

#### GET `/html`
Renderiza uma página HTML usando templates.

**Resposta:** Página HTML com design moderno

#### GET `/xml`
Retorna dados em formato XML.

**Resposta:**
```xml
<User>
  <ID>1</ID>
  <Name>João Silva</Name>
  <Email>joao@exemplo.com</Email>
  <Age>30</Age>
  <Created>2024-01-15T10:30:00Z</Created>
</User>
```

### 3. Data Binding

#### POST `/users`
Cria um novo usuário com data binding automático.

**Corpo da requisição:**
```json
{
  "name": "Maria Silva",
  "email": "maria@exemplo.com",
  "age": 25
}
```

**Resposta:**
```json
{
  "success": true,
  "message": "Usuário criado com sucesso",
  "data": {
    "id": 123,
    "name": "Maria Silva",
    "email": "maria@exemplo.com",
    "age": 25,
    "created": "2024-01-15T10:30:00Z"
  }
}
```

### 4. Query Parameters

#### GET `/search`
Demonstra o uso de query parameters.

**Parâmetros:**
- `q` (query): Termo de busca
- `limit` (query): Limite de resultados

**Exemplo:**
```bash
curl "http://localhost:8080/api/v1/search?q=laptop&limit=10"
```

**Resposta:**
```json
{
  "success": true,
  "message": "Busca realizada",
  "data": {
    "query": "laptop",
    "limit": "10",
    "results": ["resultado 1", "resultado 2", "resultado 3"]
  }
}
```

### 5. Upload e Download de Arquivos

#### POST `/upload`
Faz upload de um arquivo.

**Corpo da requisição:** `multipart/form-data`
- `file`: Arquivo a ser enviado

**Exemplo:**
```bash
curl -X POST -F "file=@documento.pdf" http://localhost:8080/api/v1/upload
```

**Resposta:**
```json
{
  "success": true,
  "message": "Arquivo enviado com sucesso",
  "data": {
    "filename": "documento.pdf",
    "size": 1024000
  }
}
```

#### GET `/download/:filename`
Faz download de um arquivo.

**Parâmetros:**
- `filename` (path): Nome do arquivo

**Exemplo:**
```bash
curl -O http://localhost:8080/api/v1/download/documento.pdf
```

### 6. Autenticação JWT

#### GET `/protected/profile`
Endpoint protegido que requer autenticação.

**Headers:**
```
Authorization: Bearer valid-token
```

**Exemplo:**
```bash
curl -H "Authorization: Bearer valid-token" http://localhost:8080/api/v1/protected/profile
```

**Resposta:**
```json
{
  "success": true,
  "message": "Perfil do usuário autenticado",
  "data": {
    "id": 1,
    "name": "Usuário Autenticado",
    "email": "user@exemplo.com",
    "age": 25,
    "created": "2024-01-15T10:30:00Z"
  }
}
```

### 7. CRUD Completo - Produtos

#### GET `/products`
Lista todos os produtos.

**Resposta:**
```json
{
  "success": true,
  "message": "Produtos listados com sucesso",
  "data": [
    {
      "id": 1,
      "name": "Laptop",
      "price": 2999.99,
      "description": "Laptop de alta performance",
      "category": "Eletrônicos"
    },
    {
      "id": 2,
      "name": "Mouse",
      "price": 89.99,
      "description": "Mouse sem fio",
      "category": "Acessórios"
    }
  ]
}
```

#### GET `/products/:id`
Obtém um produto específico.

**Parâmetros:**
- `id` (path): ID do produto

#### POST `/products`
Cria um novo produto.

**Corpo da requisição:**
```json
{
  "name": "Teclado Mecânico",
  "price": 299.99,
  "description": "Teclado mecânico RGB",
  "category": "Acessórios"
}
```

#### PUT `/products/:id`
Atualiza um produto existente.

**Parâmetros:**
- `id` (path): ID do produto

#### DELETE `/products/:id`
Remove um produto.

**Parâmetros:**
- `id` (path): ID do produto

### 8. Streaming

#### GET `/stream`
Demonstra streaming de dados em tempo real.

**Exemplo:**
```bash
curl http://localhost:8080/api/v1/stream
```

**Resposta:** Dados enviados em chunks a cada 500ms

### 9. WebSocket (Simulado)

#### GET `/ws`
Endpoint preparado para WebSocket.

**Resposta:**
```json
{
  "success": true,
  "message": "Endpoint WebSocket (implementação completa requer upgrade)",
  "data": {
    "protocol": "WebSocket",
    "status": "Ready for upgrade"
  }
}
```

## 🔧 Funcionalidades Demonstradas

### 1. **Router Otimizado**
- Roteamento de alta performance sem alocação dinâmica de memória
- Priorização inteligente de rotas
- Suporte a parâmetros de path e query

### 2. **Middleware System**
- Middleware global (Logger, Recover, CORS)
- Middleware customizado para logging
- Middleware de autenticação JWT
- Aplicação de middleware em grupos de rotas

### 3. **Data Binding**
- Binding automático de JSON
- Binding de XML
- Binding de form-data
- Validação customizada

### 4. **Data Rendering**
- Respostas JSON
- Respostas XML
- Respostas HTML (templates)
- Respostas de texto simples
- Streaming de dados
- Upload/Download de arquivos

### 5. **Templates**
- Renderização de templates HTML
- Template engine customizável
- Design responsivo e moderno

### 6. **Extensibilidade**
- Tratamento de erros centralizado
- Validador customizado
- API extensível
- Middleware customizado

### 7. **Segurança**
- Middleware CORS
- Autenticação JWT
- Preparado para TLS automático
- Suporte a HTTP/2

## 📝 Exemplos de Uso

### Testando com cURL

```bash
# Informações gerais
curl http://localhost:8080/api/v1/

# Criar usuário
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"João","email":"joao@exemplo.com","age":30}' \
  http://localhost:8080/api/v1/users

# Listar produtos
curl http://localhost:8080/api/v1/products

# Upload de arquivo
curl -X POST -F "file=@teste.txt" \
  http://localhost:8080/api/v1/upload

# Endpoint protegido
curl -H "Authorization: Bearer valid-token" \
  http://localhost:8080/api/v1/protected/profile
```

### Testando com JavaScript

```javascript
// Criar usuário
fetch('http://localhost:8080/api/v1/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    name: 'Maria',
    email: 'maria@exemplo.com',
    age: 25
  })
})
.then(response => response.json())
.then(data => console.log(data));

// Listar produtos
fetch('http://localhost:8080/api/v1/products')
.then(response => response.json())
.then(data => console.log(data));
```

## 🎯 Próximos Passos

Para expandir este playground, você pode:

1. **Implementar banco de dados real** (PostgreSQL, MongoDB)
2. **Adicionar mais validações** com bibliotecas como `go-playground/validator`
3. **Implementar WebSocket real** para comunicação em tempo real
4. **Configurar TLS automático** com Let's Encrypt
5. **Adicionar testes unitários** e de integração
6. **Implementar rate limiting** e outras medidas de segurança
7. **Adicionar documentação Swagger/OpenAPI**

## 📄 Licença

Este projeto é um playground educacional para demonstrar as funcionalidades do Echo framework.
