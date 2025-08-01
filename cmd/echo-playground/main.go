package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"echo-playground/internal"
	custommiddleware "echo-playground/pkg/middleware"
)

func main() {
	// Criar instância do Echo
	e := echo.New()

	// Configurar logger
	e.Logger.SetLevel(1) // Debug level

	// Middleware global
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())
	e.Use(custommiddleware.CustomLogger())

	// Configurar templates
	t := internal.NewTemplate()
	e.Renderer = t

	// Configurar validação de dados
	// e.Validator = &utils.CustomValidator{} // Comentado temporariamente

	// Configurar tratamento de erros centralizado
	e.HTTPErrorHandler = internal.CustomErrorHandler

	// Criar handlers
	handlers := internal.NewHandlers()
	productHandlers := internal.NewProductHandlers()

	// Grupo de rotas públicas
	public := e.Group("/api/v1")

	// Rotas básicas
	public.GET("/", handlers.HomeHandler)
	public.GET("/hello/:name", handlers.HelloHandler)
	public.GET("/html", handlers.HTMLHandler)
	public.GET("/xml", handlers.XMLHandler)

	// Documentação Swagger
	public.GET("/docs", handlers.SwaggerHandler)
	public.File("/swagger.yaml", "api/swagger.yaml")
	public.Static("/api-docs", "api")

	// Demonstração de data binding
	public.POST("/users", handlers.CreateUserHandler)

	// Endpoint de login para gerar token JWT
	public.POST("/login", handlers.LoginHandler)

	// Demonstração de query parameters
	public.GET("/search", handlers.SearchHandler)

	// Demonstração de upload de arquivo
	public.POST("/upload", handlers.UploadHandler)

	// Demonstração de download de arquivo
	public.GET("/download/:filename", handlers.DownloadHandler)

	// Demonstração de streaming
	public.GET("/stream", handlers.StreamHandler)

	// Demonstração de WebSocket (simulado)
	public.GET("/ws", handlers.WebSocketHandler)

	// Grupo de rotas protegidas (com autenticação)
	protected := e.Group("/api/v1/protected")
	protected.Use(custommiddleware.AuthMiddleware())

	protected.GET("/profile", handlers.ProfileHandler)

	// Demonstração de CRUD completo
	products := e.Group("/api/v1/products")

	// Listar produtos
	products.GET("", productHandlers.ListProductsHandler)

	// Obter produto específico
	products.GET("/:id", productHandlers.GetProductHandler)

	// Criar produto
	products.POST("", productHandlers.CreateProductHandler)

	// Atualizar produto
	products.PUT("/:id", productHandlers.UpdateProductHandler)

	// Deletar produto
	products.DELETE("/:id", productHandlers.DeleteProductHandler)

	// Obter porta da variável de ambiente ou usar padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configurar servidor HTTP/2
	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Iniciar servidor
	log.Printf("🚀 Echo Playground iniciado na porta %s", port)
	log.Printf("📖 Documentação disponível em: http://localhost:%s/api/v1/", port)
	log.Println("🔧 Recursos demonstrados:")
	log.Println("   - Router otimizado")
	log.Println("   - Middleware customizado")
	log.Println("   - Data binding")
	log.Println("   - Múltiplos formatos de resposta")
	log.Println("   - Upload/Download de arquivos")
	log.Println("   - Autenticação JWT")
	log.Println("   - CRUD completo")
	log.Println("   - Streaming")
	log.Println("   - Templates HTML")
	log.Println("   - Tratamento de erros centralizado")
	log.Println("   - Arquitetura modular Go")

	if err := e.StartServer(server); err != nil {
		log.Fatal(err)
	}
}
