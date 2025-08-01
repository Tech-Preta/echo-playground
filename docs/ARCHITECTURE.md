# 🏗️ Arquitetura do Echo Playground

Este documento descreve a arquitetura modular do Echo Playground, seguindo as melhores práticas de layout de projetos Go.

## 📋 Visão Geral

O projeto foi reestruturado seguindo o padrão de layout de projetos Go recomendado pela comunidade, proporcionando:

- **Separação clara de responsabilidades**
- **Código reutilizável e testável**
- **Organização profissional**
- **Escalabilidade para projetos maiores**

## 🏛️ Estrutura de Diretórios

### `/cmd`
**Aplicações principais do projeto**

```
cmd/
└── echo-playground/
    └── main.go          # Ponto de entrada da aplicação
```

- Contém apenas o código de inicialização
- Importa e invoca código dos diretórios `/internal` e `/pkg`
- Não contém lógica de negócio

### `/internal`
**Código privado da aplicação**

```
internal/
├── app/                 # Lógica da aplicação
│   ├── handlers.go      # Handlers HTTP
│   ├── products.go      # Handlers específicos de produtos
│   ├── template.go      # Renderizador de templates
│   └── error_handler.go # Tratamento de erros centralizado
└── pkg/                 # Pacotes internos compartilhados
```

- Código que não deve ser importado por outros projetos
- Enforçado pelo compilador Go
- Lógica específica da aplicação

### `/pkg`
**Código público reutilizável**

```
pkg/
├── api/                 # Respostas da API
│   └── response.go      # Estruturas de resposta padrão
├── middleware/          # Middlewares customizados
│   ├── auth_middleware.go    # Middleware de autenticação
│   └── logger_middleware.go  # Middleware de logging
├── models/              # Modelos de dados
│   ├── user.go          # Modelo de usuário
│   └── product.go       # Modelo de produto
└── utils/               # Utilitários
    └── validator.go     # Validador customizado
```

- Código que pode ser importado por outros projetos
- Comunica explicitamente que o código é seguro para uso
- Funcionalidades reutilizáveis

### `/web`
**Componentes específicos da web**

```
web/
└── templates/           # Templates HTML
    └── index.html       # Template principal
```

- Ativos estáticos da web
- Templates do lado do servidor
- Componentes específicos da interface

### `/configs`
**Arquivos de configuração**

```
configs/
└── config.yaml         # Configuração da aplicação
```

- Modelos de arquivo de configuração
- Configurações padrão
- Templates de configuração

### `/scripts`
**Scripts de build e deploy**

```
scripts/
├── build.sh            # Script de build
└── test_api.sh         # Script de testes da API
```

- Scripts para operações de construção
- Scripts de instalação
- Scripts de análise

### `/build`
**Configurações de build**

```
build/
└── ci/                 # CI/CD
    └── .gitkeep        # Mantém diretório no controle de versão
```

- Configurações de pacote
- Scripts de CI/CD
- Configurações de contêiner

### `/docs`
**Documentação**

```
docs/
├── API_DOCUMENTATION.md    # Documentação da API
├── FEATURES_SUMMARY.md     # Resumo das funcionalidades
└── ARCHITECTURE.md         # Este arquivo
```

- Documentação do projeto
- Documentação do usuário
- Guias de uso

### `/test`
**Testes e dados de teste**

```
test/
└── teste.txt           # Arquivo de teste para upload
```

- Aplicações de testes externos
- Dados de teste
- Testes de integração

## 🔄 Fluxo de Dados

```
Request → Middleware → Handler → Model → Response
```

1. **Request**: Requisição HTTP chega no servidor
2. **Middleware**: Processamento global (logging, CORS, auth)
3. **Handler**: Lógica específica do endpoint
4. **Model**: Manipulação de dados
5. **Response**: Resposta formatada

## 📦 Pacotes e Responsabilidades

### `cmd/echo-playground`
- **Responsabilidade**: Ponto de entrada da aplicação
- **Conteúdo**: Configuração do servidor, inicialização
- **Dependências**: Importa pacotes internos e externos

### `internal/app`
- **Responsabilidade**: Lógica da aplicação
- **Conteúdo**: Handlers, templates, tratamento de erros
- **Dependências**: Usa pacotes `/pkg` e `/internal`

### `pkg/api`
- **Responsabilidade**: Estruturas de resposta da API
- **Conteúdo**: Tipos de resposta padrão
- **Dependências**: Apenas bibliotecas padrão

### `pkg/middleware`
- **Responsabilidade**: Middlewares customizados
- **Conteúdo**: Autenticação, logging customizado
- **Dependências**: Echo framework

### `pkg/models`
- **Responsabilidade**: Modelos de dados
- **Conteúdo**: Estruturas de dados e métodos
- **Dependências**: Apenas bibliotecas padrão

### `pkg/utils`
- **Responsabilidade**: Utilitários reutilizáveis
- **Conteúdo**: Validadores, helpers
- **Dependências**: Apenas bibliotecas padrão

## 🔧 Configuração

### Arquivo de Configuração
```yaml
# configs/config.yaml
server:
  port: 8080
  host: "0.0.0.0"
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 120s
```

### Variáveis de Ambiente
```bash
# Configurações do servidor
PORT=8080
HOST=0.0.0.0

# Configurações de logging
LOG_LEVEL=debug
LOG_FORMAT=json

# Configurações de segurança
JWT_SECRET=your-secret-key
```

## 🧪 Testes

### Estrutura de Testes
```
test/
├── unit/              # Testes unitários
├── integration/       # Testes de integração
└── data/             # Dados de teste
```

### Executando Testes
```bash
# Testes unitários
go test ./...

# Testes específicos
go test ./pkg/models
go test ./internal/app

# Testes com cobertura
go test -cover ./...
```

## 🚀 Build e Deploy

### Script de Build
```bash
# Compilar aplicação
./scripts/build.sh

# Executar binário
./bin/echo-playground
```

### Docker
```dockerfile
# Dockerfile para containerização
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o echo-playground ./cmd/echo-playground

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/echo-playground .
CMD ["./echo-playground"]
```

## 📈 Benefícios da Arquitetura

### 1. **Separação de Responsabilidades**
- Cada pacote tem uma responsabilidade específica
- Fácil localização de código
- Manutenção simplificada

### 2. **Reutilização de Código**
- Pacotes `/pkg` podem ser importados por outros projetos
- Funcionalidades compartilhadas
- Redução de duplicação

### 3. **Testabilidade**
- Código modular facilita testes unitários
- Isolamento de dependências
- Cobertura de testes melhorada

### 4. **Escalabilidade**
- Estrutura preparada para crescimento
- Fácil adição de novos módulos
- Organização profissional

### 5. **Manutenibilidade**
- Código bem organizado
- Documentação clara
- Padrões consistentes

## 🔮 Próximos Passos

### Melhorias Planejadas
1. **Implementar injeção de dependência**
2. **Adicionar camada de repositório**
3. **Implementar cache distribuído**
4. **Configurar monitoramento**
5. **Adicionar métricas de performance**

### Expansão da Arquitetura
1. **Microserviços**: Dividir em serviços menores
2. **Eventos**: Implementar arquitetura orientada a eventos
3. **CQRS**: Separar comandos e consultas
4. **DDD**: Implementar Domain-Driven Design

## 📚 Referências

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Echo Framework Documentation](https://echo.labstack.com/guide/)

---

Esta arquitetura segue as melhores práticas da comunidade Go e proporciona uma base sólida para o crescimento do projeto.
