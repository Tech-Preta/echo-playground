# Exemplos JSON da API

Este diretório contém exemplos de JSON para facilitar o teste e uso da API do Echo Playground.

## Estrutura dos Arquivos

### Inputs (Entradas)

- **`create_user.json`** - Exemplo para criar um usuário
- **`create_product.json`** - Exemplo para criar um produto
- **`update_product.json`** - Exemplo para atualizar um produto
- **`search_query.json`** - Exemplo de parâmetros de busca

### Outputs (Saídas)

- **`api_response_success.json`** - Exemplo de resposta de sucesso
- **`api_response_error.json`** - Exemplo de resposta de erro
- **`products_list.json`** - Exemplo de lista de produtos
- **`users_list.json`** - Exemplo de lista de usuários
- **`upload_response.json`** - Exemplo de resposta de upload
- **`profile_response.json`** - Exemplo de resposta de perfil autenticado

## Como Usar

### Testando com cURL

```bash
# Criar usuário
curl -X POST -H "Content-Type: application/json" \
  -d @examples/create_user.json \
  http://localhost:8080/api/v1/users

# Criar produto
curl -X POST -H "Content-Type: application/json" \
  -d @examples/create_product.json \
  http://localhost:8080/api/v1/products

# Atualizar produto
curl -X PUT -H "Content-Type: application/json" \
  -d @examples/update_product.json \
  http://localhost:8080/api/v1/products/1
```

### Testando com JavaScript

```javascript
// Criar usuário
const userData = require('./examples/create_user.json');

fetch('http://localhost:8080/api/v1/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(userData)
})
.then(response => response.json())
.then(data => console.log(data));
```

### Testando com Python

```python
import json
import requests

# Carregar dados do exemplo
with open('examples/create_user.json', 'r') as f:
    user_data = json.load(f)

# Fazer requisição
response = requests.post(
    "http://localhost:8080/api/v1/users",
    json=user_data
)
print(response.json())
```

## Estrutura dos Exemplos

### Inputs

Todos os inputs seguem a estrutura definida no Swagger (`api/swagger.yaml`):

- **Usuários**: `name`, `email`, `age`
- **Produtos**: `name`, `price`, `description`, `category`
- **Busca**: `query`, `limit`, `category`, `min_price`, `max_price`

### Outputs

Todos os outputs seguem o padrão da API:

```json
{
  "success": true/false,
  "message": "Mensagem descritiva",
  "data": { ... },
  "error": "Detalhes do erro (opcional)"
}
```

## Validação

Os exemplos são validados contra:

1. **Swagger Schema** (`api/swagger.yaml`)
2. **Go Structs** (`pkg/models/`)
3. **API Response Format** (`pkg/api/response.go`)

## Contribuindo

Ao adicionar novos exemplos:

1. Mantenha a consistência com os schemas existentes
2. Use dados realistas mas fictícios
3. Inclua comentários explicativos quando necessário
4. Atualize este README se necessário
