package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"echo-playground/pkg/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Handlers contém todos os handlers da aplicação
type Handlers struct{}

// NewHandlers cria uma nova instância de handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

// HomeHandler retorna informações sobre o framework
func (h *Handlers) HomeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Bem-vindo ao Echo Playground!",
		"data": map[string]interface{}{
			"framework": "Echo",
			"version":   "v4",
			"features": []string{
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
		},
	})
}

// HelloHandler retorna uma saudação personalizada
func (h *Handlers) HelloHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, fmt.Sprintf("Olá, %s! Bem-vindo ao Echo!", name))
}

// HTMLHandler renderiza uma página HTML
func (h *Handlers) HTMLHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title":   "Echo Playground",
		"Message": "Esta é uma página HTML renderizada pelo Echo!",
	})
}

// XMLHandler retorna dados em formato XML
func (h *Handlers) XMLHandler(c echo.Context) error {
	user := models.NewUser("João Silva", "joao@exemplo.com", 30)
	user.SetID(1)
	return c.XML(http.StatusOK, user)
}

// CreateUserHandler cria um novo usuário
func (h *Handlers) CreateUserHandler(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Erro ao processar dados",
			"error":   err.Error(),
		})
	}

	// Simular validação
	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Nome é obrigatório",
			"error":   "",
		})
	}

	user.SetID(123) // Simular ID gerado
	user.Created = time.Now().Format(time.RFC3339)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Usuário criado com sucesso",
		"data":    user,
	})
}

// SearchHandler demonstra query parameters
func (h *Handlers) SearchHandler(c echo.Context) error {
	query := c.QueryParam("q")
	limit := c.QueryParam("limit")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Busca realizada",
		"data": map[string]interface{}{
			"query":   query,
			"limit":   limit,
			"results": []string{"resultado 1", "resultado 2", "resultado 3"},
		},
	})
}

// UploadHandler faz upload de arquivo
func (h *Handlers) UploadHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Erro ao processar arquivo",
			"error":   err.Error(),
		})
	}

	// Salvar arquivo
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Erro ao abrir arquivo",
			"error":   "",
		})
	}
	defer func() {
		if err := src.Close(); err != nil {
			c.Logger().Error("Erro ao fechar arquivo fonte:", err)
		}
	}()

	dst, err := os.Create("uploads/" + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Erro ao salvar arquivo",
			"error":   err.Error(),
		})
	}
	defer func() {
		if err := dst.Close(); err != nil {
			c.Logger().Error("Erro ao fechar arquivo destino:", err)
		}
	}()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Erro ao copiar arquivo",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Arquivo enviado com sucesso",
		"data": map[string]interface{}{
			"filename": file.Filename,
			"size":     file.Size,
		},
	})
}

// DownloadHandler faz download de arquivo
func (h *Handlers) DownloadHandler(c echo.Context) error {
	filename := c.Param("filename")
	return c.Attachment("uploads/"+filename, filename)
}

// ProfileHandler retorna perfil do usuário autenticado
func (h *Handlers) ProfileHandler(c echo.Context) error {
	user := models.NewUser("Usuário Autenticado", "user@exemplo.com", 25)
	user.SetID(1)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Perfil do usuário autenticado",
		"data":    user,
	})
}

// StreamHandler demonstra streaming
func (h *Handlers) StreamHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/plain")
	c.Response().Header().Set("Transfer-Encoding", "chunked")

	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		if _, err := fmt.Fprintf(c.Response().Writer, "Chunk %d\n", i); err != nil {
			return err
		}
		c.Response().Flush()
	}

	return nil
}

// WebSocketHandler demonstra WebSocket (simulado)
func (h *Handlers) WebSocketHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Endpoint WebSocket (implementação completa requer upgrade)",
		"data": map[string]interface{}{
			"protocol": "WebSocket",
			"status":   "Ready for upgrade",
		},
	})
}

// LoginHandler gera um token JWT para teste
func (h *Handlers) LoginHandler(c echo.Context) error {
	// Para demonstração, vamos aceitar qualquer usuário/senha
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	loginReq := new(LoginRequest)
	if err := c.Bind(loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Dados de login inválidos",
			"error":   err.Error(),
		})
	}

	// Validação simples para demonstração
	if loginReq.Username == "" || loginReq.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Username e password são obrigatórios",
			"error":   "",
		})
	}

	// Claims do token
	claims := jwt.MapClaims{
		"user_id":  123,
		"username": loginReq.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Expira em 72 horas
	}

	// Criar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("secret") // Mesma secret usada no middleware
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Erro ao gerar token",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Login realizado com sucesso",
		"data": map[string]interface{}{
			"token":    tokenString,
			"user_id":  123,
			"username": loginReq.Username,
			"expires":  time.Now().Add(time.Hour * 72).Unix(),
		},
	})
}

// SwaggerHandler serve a documentação Swagger
func (h *Handlers) SwaggerHandler(c echo.Context) error {
	swaggerHTML := `
<!DOCTYPE html>
<html>
<head>
    <title>Echo Playground - API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui.css" />
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }
        *, *:before, *:after {
            box-sizing: inherit;
        }
        body {
            margin:0;
            background: #fafafa;
        }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '/api/v1/swagger.yaml',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
        };
    </script>
</body>
</html>`
	return c.HTML(http.StatusOK, swaggerHTML)
}
