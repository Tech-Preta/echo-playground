# 🚀 Echo Playground

Um playground completo que demonstra todas as funcionalidades do **Echo framework** - um framework Go de alta performance, extensível e minimalista, seguindo as melhores práticas de layout de projetos Go.

## ✨ Funcionalidades Demonstradas

### 🔥 **High Performance Router**
- Router otimizado sem alocação dinâmica de memória
- Priorização inteligente de rotas
- Processamento eficiente de requisições HTTP

### 📈 **Scalable REST APIs**
- Organização de endpoints em grupos lógicos
- APIs RESTful robustas e escaláveis
- Gerenciamento simplificado de APIs complexas

### 🔒 **Automatic TLS**
- Configuração preparada para TLS automático
- Integração com Let's Encrypt
- Comunicação segura simplificada

### ⚡ **HTTP/2 Support**
- Suporte nativo ao protocolo HTTP/2
- Melhoria na velocidade e responsividade
- Otimização na transmissão de dados

### 🔧 **Middleware System**
- Middleware global (Logger, Recover, CORS)
- Middleware customizado para logging
- Middleware de autenticação JWT
- Aplicação flexível em diferentes níveis

### 📊 **Data Binding**
- Binding automático de JSON, XML e form-data
- Extração simplificada de dados de requisições
- Integração perfeita com lógica de aplicação

### 🎨 **Data Rendering**
- Respostas em múltiplos formatos (JSON, XML, HTML)
- Upload e download de arquivos
- Streaming de dados em tempo real
- Templates HTML dinâmicos

### 🎯 **Template Support**
- Renderização com qualquer engine de template
- Templates HTML dinâmicos
- Design responsivo e moderno

### 🔌 **Extensible Architecture**
- Tratamento de erros centralizado
- API facilmente extensível
- Middleware customizado
- Integração com funcionalidades de terceiros

### 🏗️ **Arquitetura Modular Go**
- Estrutura seguindo padrões da comunidade Go
- Separação clara de responsabilidades
- Código reutilizável e testável
- Organização profissional

## 🚀 Início Rápido

### Pré-requisitos
- Go 1.21 ou superior
- Git

### Instalação

```bash
# Clonar o repositório
git clone https://github.com/nataliagranato/echo.git
cd echo

# Instalar dependências
go mod tidy

# Executar o servidor
go run cmd/echo-playground/main.go
```

O servidor estará disponível em: `http://localhost:8080`

### Build da Aplicação

```bash
# Compilar a aplicação
./scripts/build.sh

# Executar o binário
./bin/echo-playground
```

## 📚 Endpoints Principais

### Informações Gerais
- `GET /api/v1/` - Informações sobre o framework

### Demonstrações
- `GET /api/v1/hello/:name` - Saudação personalizada
- `GET /api/v1/html` - Página HTML renderizada
- `GET /api/v1/xml` - Resposta em XML
- `GET /api/v1/stream` - Streaming de dados

### Data Binding
- `POST /api/v1/users` - Criar usuário com binding automático
- `GET /api/v1/search` - Query parameters

### Upload/Download
- `POST /api/v1/upload` - Upload de arquivos
- `GET /api/v1/download/:filename` - Download de arquivos

### Autenticação
- `GET /api/v1/protected/profile` - Endpoint protegido (requer JWT)

### CRUD Completo
- `GET /api/v1/products` - Listar produtos
- `GET /api/v1/products/:id` - Obter produto
- `POST /api/v1/products` - Criar produto
- `PUT /api/v1/products/:id` - Atualizar produto
- `DELETE /api/v1/products/:id` - Deletar produto

## 🏗️ Estrutura do Projeto

```
echo-playground/
├── cmd/                    # Aplicações principais
│   └── echo-playground/   # Aplicação principal
│       └── main.go        # Ponto de entrada
├── internal/              # Código privado da aplicação
│   ├── app/              # Lógica da aplicação
│   │   ├── handlers.go   # Handlers HTTP
│   │   ├── products.go   # Handlers de produtos
│   │   ├── template.go   # Renderizador de templates
│   │   └── error_handler.go # Tratamento de erros
│   └── pkg/              # Pacotes internos compartilhados
├── pkg/                   # Código público reutilizável
│   ├── api/              # Respostas da API
│   ├── middleware/       # Middlewares customizados
│   ├── models/           # Modelos de dados
│   └── utils/            # Utilitários
├── web/                  # Componentes web
│   └── templates/        # Templates HTML
├── configs/              # Arquivos de configuração
├── scripts/              # Scripts de build e deploy
├── build/                # Configurações de build
│   └── ci/              # CI/CD
├── docs/                 # Documentação
├── test/                 # Testes e dados de teste
├── examples/             # Exemplos de uso
├── assets/               # Recursos estáticos
└── go.mod               # Dependências do Go
```

## 🧪 Testando a API

### Exemplos com cURL

```bash
# Informações gerais
curl http://localhost:8080/api/v1/

# Criar usuário
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"João","email":"joao@exemplo.com","age":30}' \
  http://localhost:8080/api/v1/users

# Listar produtos
curl http://localhost:8080/api/v1/products

# Endpoint protegido
curl -H "Authorization: Bearer valid-token" \
  http://localhost:8080/api/v1/protected/profile

# Upload de arquivo
curl -X POST -F "file=@test/teste.txt" \
  http://localhost:8080/api/v1/upload
```

### Script de Teste Automatizado

```bash
# Executar todos os testes
./scripts/test_api.sh
```

## 📖 Documentação Completa

Para documentação detalhada de todos os endpoints e funcionalidades, consulte:
- [Documentação da API](./docs/API_DOCUMENTATION.md)
- [Resumo das Funcionalidades](./docs/FEATURES_SUMMARY.md)

## 🔧 Configurações Avançadas

### Configuração do Servidor
As configurações do servidor estão em `configs/config.yaml`:

```yaml
server:
  port: 8080
  host: "0.0.0.0"
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 120s
```

### TLS Automático
Para habilitar TLS automático com Let's Encrypt, modifique o arquivo `cmd/echo-playground/main.go`:

```go
autoTLS := autocert.Manager{
    Prompt: autocert.AcceptTOS,
    HostPolicy: autocert.HostWhitelist("seu-dominio.com"),
    Cache: autocert.DirCache("/var/www/.cache"),
}
e.AutoTLSManager = &autoTLS
```

### HTTP/2
Para habilitar HTTP/2, descomente as linhas de configuração TLS no `cmd/echo-playground/main.go`.

## 🎯 Próximos Passos

Para expandir este playground, você pode:

1. **Implementar banco de dados real** (PostgreSQL, MongoDB)
2. **Adicionar validações avançadas** com `go-playground/validator`
3. **Implementar WebSocket real** para comunicação em tempo real
4. **Configurar rate limiting** e outras medidas de segurança
5. **Adicionar testes unitários** e de integração
6. **Implementar documentação Swagger/OpenAPI**
7. **Adicionar monitoramento** e métricas
8. **Configurar CI/CD** com GitHub Actions ou GitLab CI

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para:

- Reportar bugs
- Sugerir novas funcionalidades
- Melhorar a documentação
- Adicionar exemplos de uso

## 📄 Licença

Este projeto é um playground educacional para demonstrar as funcionalidades do Echo framework.

## 🔗 Links Úteis

- [Echo Framework](https://echo.labstack.com/)
- [Documentação Oficial](https://echo.labstack.com/guide/)
- [GitHub do Echo](https://github.com/labstack/echo)
- [Go Programming Language](https://golang.org/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
