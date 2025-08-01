// Package docs fornece documentação para o Echo Playground.
//
// Este pacote contém documentação gerada automaticamente para a API
// do Echo Playground, incluindo exemplos de uso e referências.
package docs

import "time"

// EchoPlayground representa a documentação principal do projeto
type EchoPlayground struct {
	// Nome do projeto
	Name string `json:"name"`

	// Versão atual
	Version string `json:"version"`

	// Descrição do projeto
	Description string `json:"description"`

	// Funcionalidades principais
	Features []string `json:"features"`

	// Informações de contato
	Contact ContactInfo `json:"contact"`

	// Licença
	License LicenseInfo `json:"license"`

	// Data de criação da documentação
	CreatedAt time.Time `json:"created_at"`
}

// ContactInfo contém informações de contato
type ContactInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LicenseInfo contém informações sobre a licença
type LicenseInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewEchoPlayground cria uma nova instância da documentação
func NewEchoPlayground() *EchoPlayground {
	return &EchoPlayground{
		Name:        "Echo Playground",
		Version:     "1.0.0",
		Description: "Um playground completo que demonstra todas as funcionalidades do Echo framework",
		Features: []string{
			"High Performance Router",
			"Scalable REST APIs",
			"Automatic TLS",
			"HTTP/2 Support",
			"Middleware System",
			"Data Binding",
			"Data Rendering",
			"Template Support",
			"Extensible Architecture",
		},
		Contact: ContactInfo{
			Name: "Echo Playground",
			URL:  "https://echo.labstack.com",
		},
		License: LicenseInfo{
			Name: "MIT",
			URL:  "https://opensource.org/licenses/MIT",
		},
		CreatedAt: time.Now(),
	}
}

// APIEndpoints lista todos os endpoints da API
type APIEndpoints struct {
	// Endpoints de informações
	Info []Endpoint `json:"info"`

	// Endpoints de demonstrações
	Demos []Endpoint `json:"demos"`

	// Endpoints de usuários
	Users []Endpoint `json:"users"`

	// Endpoints de produtos
	Products []Endpoint `json:"products"`

	// Endpoints de upload/download
	Files []Endpoint `json:"files"`

	// Endpoints de autenticação
	Auth []Endpoint `json:"auth"`

	// Endpoints de streaming
	Streaming []Endpoint `json:"streaming"`
}

// Endpoint representa um endpoint da API
type Endpoint struct {
	Path        string      `json:"path"`
	Method      string      `json:"method"`
	Description string      `json:"description"`
	Tags        []string    `json:"tags"`
	Parameters  []Parameter `json:"parameters,omitempty"`
	Responses   []Response  `json:"responses"`
}

// Parameter representa um parâmetro de endpoint
type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Example     string `json:"example,omitempty"`
}

// Response representa uma resposta de endpoint
type Response struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	ContentType string `json:"content_type"`
	Schema      string `json:"schema,omitempty"`
}

// GetAPIEndpoints retorna todos os endpoints da API
func GetAPIEndpoints() *APIEndpoints {
	return &APIEndpoints{
		Info: []Endpoint{
			{
				Path:        "/api/v1/",
				Method:      "GET",
				Description: "Informações sobre o framework",
				Tags:        []string{"Informações"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Informações do framework",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
		},
		Demos: []Endpoint{
			{
				Path:        "/api/v1/hello/{name}",
				Method:      "GET",
				Description: "Saudação personalizada",
				Tags:        []string{"Demonstrações"},
				Parameters: []Parameter{
					{
						Name:        "name",
						In:          "path",
						Required:    true,
						Description: "Nome da pessoa",
						Type:        "string",
						Example:     "João",
					},
				},
				Responses: []Response{
					{
						Code:        200,
						Description: "Saudação personalizada",
						ContentType: "text/plain",
					},
				},
			},
			{
				Path:        "/api/v1/html",
				Method:      "GET",
				Description: "Página HTML renderizada",
				Tags:        []string{"Demonstrações"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Página HTML",
						ContentType: "text/html",
					},
				},
			},
			{
				Path:        "/api/v1/xml",
				Method:      "GET",
				Description: "Resposta em XML",
				Tags:        []string{"Demonstrações"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Dados em XML",
						ContentType: "application/xml",
						Schema:      "User",
					},
				},
			},
			{
				Path:        "/api/v1/search",
				Method:      "GET",
				Description: "Busca com query parameters",
				Tags:        []string{"Demonstrações"},
				Parameters: []Parameter{
					{
						Name:        "q",
						In:          "query",
						Required:    false,
						Description: "Termo de busca",
						Type:        "string",
						Example:     "laptop",
					},
					{
						Name:        "limit",
						In:          "query",
						Required:    false,
						Description: "Limite de resultados",
						Type:        "string",
						Example:     "10",
					},
				},
				Responses: []Response{
					{
						Code:        200,
						Description: "Resultados da busca",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
		},
		Users: []Endpoint{
			{
				Path:        "/api/v1/users",
				Method:      "POST",
				Description: "Criar usuário",
				Tags:        []string{"Usuários"},
				Responses: []Response{
					{
						Code:        201,
						Description: "Usuário criado com sucesso",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
					{
						Code:        400,
						Description: "Dados inválidos",
						ContentType: "application/json",
						Schema:      "ErrorResponse",
					},
				},
			},
		},
		Products: []Endpoint{
			{
				Path:        "/api/v1/products",
				Method:      "GET",
				Description: "Listar produtos",
				Tags:        []string{"Produtos"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Lista de produtos",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
			{
				Path:        "/api/v1/products",
				Method:      "POST",
				Description: "Criar produto",
				Tags:        []string{"Produtos"},
				Responses: []Response{
					{
						Code:        201,
						Description: "Produto criado com sucesso",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
			{
				Path:        "/api/v1/products/{id}",
				Method:      "GET",
				Description: "Obter produto",
				Tags:        []string{"Produtos"},
				Parameters: []Parameter{
					{
						Name:        "id",
						In:          "path",
						Required:    true,
						Description: "ID do produto",
						Type:        "string",
						Example:     "1",
					},
				},
				Responses: []Response{
					{
						Code:        200,
						Description: "Produto encontrado",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
			{
				Path:        "/api/v1/products/{id}",
				Method:      "PUT",
				Description: "Atualizar produto",
				Tags:        []string{"Produtos"},
				Parameters: []Parameter{
					{
						Name:        "id",
						In:          "path",
						Required:    true,
						Description: "ID do produto",
						Type:        "string",
						Example:     "1",
					},
				},
				Responses: []Response{
					{
						Code:        200,
						Description: "Produto atualizado com sucesso",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
			{
				Path:        "/api/v1/products/{id}",
				Method:      "DELETE",
				Description: "Deletar produto",
				Tags:        []string{"Produtos"},
				Parameters: []Parameter{
					{
						Name:        "id",
						In:          "path",
						Required:    true,
						Description: "ID do produto",
						Type:        "string",
						Example:     "1",
					},
				},
				Responses: []Response{
					{
						Code:        200,
						Description: "Produto deletado com sucesso",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
				},
			},
		},
		Files: []Endpoint{
			{
				Path:        "/api/v1/upload",
				Method:      "POST",
				Description: "Upload de arquivo",
				Tags:        []string{"Upload/Download"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Arquivo enviado com sucesso",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
					{
						Code:        400,
						Description: "Erro no upload",
						ContentType: "application/json",
						Schema:      "ErrorResponse",
					},
				},
			},
			{
				Path:        "/api/v1/download/{filename}",
				Method:      "GET",
				Description: "Download de arquivo",
				Tags:        []string{"Upload/Download"},
				Parameters: []Parameter{
					{
						Name:        "filename",
						In:          "path",
						Required:    true,
						Description: "Nome do arquivo",
						Type:        "string",
						Example:     "documento.pdf",
					},
				},
				Responses: []Response{
					{
						Code:        200,
						Description: "Arquivo para download",
						ContentType: "application/octet-stream",
					},
				},
			},
		},
		Auth: []Endpoint{
			{
				Path:        "/api/v1/protected/profile",
				Method:      "GET",
				Description: "Perfil do usuário autenticado",
				Tags:        []string{"Autenticação"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Perfil do usuário",
						ContentType: "application/json",
						Schema:      "APIResponse",
					},
					{
						Code:        401,
						Description: "Não autorizado",
						ContentType: "application/json",
						Schema:      "ErrorResponse",
					},
				},
			},
		},
		Streaming: []Endpoint{
			{
				Path:        "/api/v1/stream",
				Method:      "GET",
				Description: "Streaming de dados",
				Tags:        []string{"Streaming"},
				Responses: []Response{
					{
						Code:        200,
						Description: "Dados em streaming",
						ContentType: "text/plain",
					},
				},
			},
		},
	}
}

// Examples fornece exemplos de uso da API
type Examples struct {
	// Exemplos de cURL
	Curl []string `json:"curl"`

	// Exemplos de JavaScript
	JavaScript []string `json:"javascript"`

	// Exemplos de Python
	Python []string `json:"python"`
}

// GetExamples retorna exemplos de uso da API
func GetExamples() *Examples {
	return &Examples{
		Curl: []string{
			`# Informações gerais
curl http://localhost:8080/api/v1/`,

			`# Criar usuário
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Maria","email":"maria@exemplo.com","age":25}' \
  http://localhost:8080/api/v1/users`,

			`# Listar produtos
curl http://localhost:8080/api/v1/products`,

			`# Endpoint protegido
curl -H "Authorization: Bearer valid-token" \
  http://localhost:8080/api/v1/protected/profile`,

			`# Upload de arquivo
curl -X POST -F "file=@teste.txt" \
  http://localhost:8080/api/v1/upload`,
		},
		JavaScript: []string{
			`// Criar usuário
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
.then(data => console.log(data));`,

			`// Listar produtos
fetch('http://localhost:8080/api/v1/products')
.then(response => response.json())
.then(data => console.log(data));`,
		},
		Python: []string{
			`# Criar usuário
import requests

data = {
    "name": "Maria",
    "email": "maria@exemplo.com",
    "age": 25
}

response = requests.post(
    "http://localhost:8080/api/v1/users",
    json=data
)
print(response.json())`,

			`# Listar produtos
import requests

response = requests.get("http://localhost:8080/api/v1/products")
print(response.json())`,
		},
	}
}

// Architecture fornece informações sobre a arquitetura do projeto
type Architecture struct {
	// Estrutura de diretórios
	Directories []string `json:"directories"`

	// Pacotes principais
	Packages []Package `json:"packages"`

	// Fluxo de dados
	DataFlow string `json:"data_flow"`
}

// Package representa um pacote do projeto
type Package struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Responsibility string   `json:"responsibility"`
	Files          []string `json:"files"`
}

// GetArchitecture retorna informações sobre a arquitetura
func GetArchitecture() *Architecture {
	return &Architecture{
		Directories: []string{
			"cmd/echo-playground/ - Ponto de entrada da aplicação",
			"internal/app/ - Lógica da aplicação",
			"pkg/api/ - Estruturas de resposta da API",
			"pkg/middleware/ - Middlewares customizados",
			"pkg/models/ - Modelos de dados",
			"pkg/utils/ - Utilitários",
			"web/templates/ - Templates HTML",
			"configs/ - Arquivos de configuração",
			"scripts/ - Scripts de build",
			"docs/ - Documentação",
		},
		Packages: []Package{
			{
				Name:           "cmd/echo-playground",
				Description:    "Aplicação principal",
				Responsibility: "Ponto de entrada da aplicação",
				Files:          []string{"main.go"},
			},
			{
				Name:           "internal/app",
				Description:    "Lógica da aplicação",
				Responsibility: "Handlers, templates, tratamento de erros",
				Files:          []string{"handlers.go", "products.go", "template.go", "error_handler.go"},
			},
			{
				Name:           "pkg/api",
				Description:    "Respostas da API",
				Responsibility: "Estruturas de resposta padrão",
				Files:          []string{"response.go"},
			},
			{
				Name:           "pkg/middleware",
				Description:    "Middlewares customizados",
				Responsibility: "Autenticação, logging customizado",
				Files:          []string{"auth_middleware.go", "logger_middleware.go"},
			},
			{
				Name:           "pkg/models",
				Description:    "Modelos de dados",
				Responsibility: "Estruturas de dados e métodos",
				Files:          []string{"user.go", "product.go"},
			},
			{
				Name:           "pkg/utils",
				Description:    "Utilitários",
				Responsibility: "Validadores, helpers",
				Files:          []string{"validator.go"},
			},
		},
		DataFlow: "Request → Middleware → Handler → Model → Response",
	}
}
