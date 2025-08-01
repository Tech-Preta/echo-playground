# Makefile para Echo Playground
# Uso: make <comando>

.PHONY: help dev build test lint fmt vet clean setup install-tools docs-serve docs-generate api-test api-curl build-linux build-darwin build-windows clean-build watch debug test-unit test-integration test-e2e benchmark quality pre-commit docs-build docs-serve-local example-curl example-js example-python start stop restart logs

# Comandos principais
help: ## Mostra esta ajuda
	@echo "Comandos disponíveis:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Executa a aplicação em modo desenvolvimento
	go run cmd/echo-playground/main.go

build: ## Compila a aplicação
	go build -o bin/echo cmd/echo-playground/main.go

test: ## Executa todos os testes
	go test ./...

test-verbose: ## Executa testes com output verboso
	go test -v ./...

test-coverage: ## Executa testes com cobertura
	go test -cover ./...

lint: ## Executa linter
	golangci-lint run

fmt: ## Formata código
	go fmt ./...

vet: ## Executa go vet
	go vet ./...

mod-tidy: ## Organiza dependências
	go mod tidy

mod-verify: ## Verifica dependências
	go mod verify

# Scripts de documentação
docs-serve: ## Serve documentação Swagger
	swagger serve api/swagger.yaml

docs-generate: ## Gera especificação Swagger
	swagger generate spec -o api/swagger.yaml

# Scripts de API
api-test: ## Executa testes da API
	./scripts/test_api.sh

api-curl: ## Testa API com curl
	curl -s http://localhost:8080/api/v1/ | jq

# Scripts de build e deploy
build-linux: ## Compila para Linux
	GOOS=linux GOARCH=amd64 go build -o bin/echo-linux cmd/echo-playground/main.go

build-darwin: ## Compila para macOS
	GOOS=darwin GOARCH=amd64 go build -o bin/echo-darwin cmd/echo-playground/main.go

build-windows: ## Compila para Windows
	GOOS=windows GOARCH=amd64 go build -o bin/echo-windows.exe cmd/echo-playground/main.go

# Scripts de limpeza
clean: ## Remove arquivos de build
	rm -rf bin/ dist/ coverage.out

clean-build: ## Limpa e recompila
	go clean && rm -rf bin/

# Scripts de desenvolvimento local
setup: ## Configura o ambiente
	go mod download && go mod verify

install-tools: ## Instala ferramentas de desenvolvimento
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-swagger: ## Instala Swagger
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# Scripts específicos do projeto
start: ## Inicia a aplicação
	go run cmd/echo-playground/main.go

stop: ## Para a aplicação
	pkill -f echo-playground

restart: stop start ## Reinicia a aplicação

logs: ## Mostra logs da aplicação
	tail -f logs/echo-playground.log

# Scripts de desenvolvimento
watch: ## Hot reload com Air
	air -c .air.toml

debug: ## Debug com Delve
	dlv debug cmd/echo-playground/main.go

# Scripts de teste
test-unit: ## Testes unitários
	go test -short ./...

test-integration: ## Testes de integração
	go test -tags=integration ./...

test-e2e: ## Testes end-to-end
	go test -tags=e2e ./...

benchmark: ## Benchmarks
	go test -bench=. ./...

# Scripts de qualidade
quality: lint test fmt ## Executa lint, test e fmt

pre-commit: quality build ## Executa qualidade e build

# Scripts de documentação
docs-build: ## Gera e valida documentação Swagger
	swagger generate spec -o api/swagger.yaml && swagger validate api/swagger.yaml

docs-serve-local: ## Serve documentação local
	swagger serve -F=swagger api/swagger.yaml --port=8081

# Scripts de exemplo
example-curl: ## Testa com exemplo curl
	curl -X POST -H 'Content-Type: application/json' -d @examples/create_user.json http://localhost:8080/api/v1/users

example-js: ## Executa exemplo JavaScript
	node examples/test.js

example-python: ## Executa exemplo Python
	python examples/test.py

# Scripts de teste com exemplos
test-with-examples: ## Testa API com exemplos JSON
	./scripts/test_with_examples.sh

# Scripts de Docker
docker-build: ## Build Docker image
	docker build -t echo-playground .

docker-run: ## Executa container Docker
	docker run -p 8080:8080 echo-playground

# Scripts de deploy
deploy-dev: build ## Deploy para desenvolvimento
	@echo "Deploying to development..."

deploy-prod: build ## Deploy para produção
	@echo "Deploying to production..."

# Scripts de monitoramento
health-check: ## Verifica saúde da aplicação
	curl -f http://localhost:8080/api/v1/ || exit 1

# Scripts de backup
backup: ## Backup dos dados
	@echo "Creating backup..."

# Scripts de restore
restore: ## Restore dos dados
	@echo "Restoring from backup..."

# Scripts de migração
migrate: ## Executa migrações
	@echo "Running migrations..."

# Scripts de seed
seed: ## Popula banco com dados de teste
	@echo "Seeding database..."

# Scripts de análise
analyze: ## Análise estática do código
	go vet ./...
	golangci-lint run
	gosec ./...

# Scripts de performance
profile: ## Gera perfil de performance
	go test -cpuprofile=cpu.prof -memprofile=mem.prof ./...

# Scripts de segurança
security-scan: ## Scan de segurança
	gosec ./...

# Scripts de documentação completa
docs-all: docs-build docs-serve-local ## Gera e serve toda documentação

# Scripts de teste completo
test-all: test-unit test-integration test-e2e benchmark ## Executa todos os tipos de teste

# Scripts de build completo
build-all: build-linux build-darwin build-windows ## Compila para todas as plataformas

# Scripts de desenvolvimento completo
dev-setup: setup install-tools install-swagger ## Configuração completa do ambiente de desenvolvimento
