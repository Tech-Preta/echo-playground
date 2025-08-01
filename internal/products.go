package internal

import (
	"fmt"
	"net/http"

	"echo-playground/pkg/models"

	"github.com/labstack/echo/v4"
)

// ProductHandlers contém os handlers relacionados a produtos
type ProductHandlers struct{}

// NewProductHandlers cria uma nova instância de handlers de produtos
func NewProductHandlers() *ProductHandlers {
	return &ProductHandlers{}
}

// ListProductsHandler lista todos os produtos
func (h *ProductHandlers) ListProductsHandler(c echo.Context) error {
	products := []*models.Product{
		models.NewProduct("Laptop", "Laptop de alta performance", "Eletrônicos", 2999.99),
		models.NewProduct("Mouse", "Mouse sem fio", "Acessórios", 89.99),
		models.NewProduct("Teclado", "Teclado mecânico", "Acessórios", 199.99),
	}

	// Definir IDs
	products[0].SetID(1)
	products[1].SetID(2)
	products[2].SetID(3)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Produtos listados com sucesso",
		"data":    products,
	})
}

// GetProductHandler obtém um produto específico
func (h *ProductHandlers) GetProductHandler(c echo.Context) error {
	_ = c.Param("id") // Usar o parâmetro para evitar warning
	// Simular busca no banco de dados
	product := models.NewProduct("Laptop", "Laptop de alta performance", "Eletrônicos", 2999.99)
	product.SetID(1)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Produto encontrado",
		"data":    product,
	})
}

// CreateProductHandler cria um novo produto
func (h *ProductHandlers) CreateProductHandler(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Erro ao processar dados do produto",
			"error":   "",
		})
	}

	product.SetID(4) // Simular ID gerado

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Produto criado com sucesso",
		"data":    product,
	})
}

// UpdateProductHandler atualiza um produto existente
func (h *ProductHandlers) UpdateProductHandler(c echo.Context) error {
	_ = c.Param("id") // Usar o parâmetro para evitar warning
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Erro ao processar dados do produto",
			"error":   "",
		})
	}

	product.SetID(1) // Simular ID do produto atualizado

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Produto atualizado com sucesso",
		"data":    product,
	})
}

// DeleteProductHandler remove um produto
func (h *ProductHandlers) DeleteProductHandler(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Produto com ID %s deletado com sucesso", id),
	})
}
