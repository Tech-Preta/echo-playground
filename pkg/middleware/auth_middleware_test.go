package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	}

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := AuthMiddleware()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	if strings.TrimSpace(rec.Body.String()) != "success" {
		t.Errorf("Expected 'success', got '%s'", strings.TrimSpace(rec.Body.String()))
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	}

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := AuthMiddleware()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// O middleware retorna JSON response, não erro HTTP
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}

	// Verificar se a resposta contém mensagem de erro
	responseBody := rec.Body.String()
	if !strings.Contains(responseBody, "Token inválido") {
		t.Errorf("Expected response to contain 'Token inválido', got '%s'", responseBody)
	}
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	}

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := AuthMiddleware()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// O middleware retorna JSON response, não erro HTTP
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}

	// Verificar se a resposta contém mensagem de erro
	responseBody := rec.Body.String()
	if !strings.Contains(responseBody, "Token de autorização não fornecido") {
		t.Errorf("Expected response to contain 'Token de autorização não fornecido', got '%s'", responseBody)
	}
}

func TestAuthMiddleware_InvalidAuthorizationFormat(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	}

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "InvalidFormat token")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := AuthMiddleware()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// O middleware retorna JSON response, não erro HTTP
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}

	// Verificar se a resposta contém mensagem de erro
	responseBody := rec.Body.String()
	if !strings.Contains(responseBody, "Token inválido") {
		t.Errorf("Expected response to contain 'Token inválido', got '%s'", responseBody)
	}
}
