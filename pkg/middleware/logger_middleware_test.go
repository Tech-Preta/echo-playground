package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestLoggerMiddleware(t *testing.T) {
	e := echo.New()

	// Capturar a saída de log
	var buf bytes.Buffer

	// Criar um pipe para capturar stdout
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	}

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := CustomLogger()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Restaurar stdout
	if err := w.Close(); err != nil {
		t.Errorf("Failed to close writer: %v", err)
	}
	os.Stdout = stdout

	// Ler a saída capturada
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("Failed to copy output: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Errorf("Failed to close reader: %v", err)
	}

	// Verificar se a resposta foi bem sucedida
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	if strings.TrimSpace(rec.Body.String()) != "success" {
		t.Errorf("Expected 'success', got '%s'", strings.TrimSpace(rec.Body.String()))
	}
}

func TestLoggerMiddleware_POSTRequest(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
	}

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(`{"name":"test"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := CustomLogger()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verificar se a resposta foi bem sucedida
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}

	responseBody := strings.TrimSpace(rec.Body.String())
	if !strings.Contains(responseBody, "created") {
		t.Errorf("Expected response to contain 'created', got '%s'", responseBody)
	}
}

func TestLoggerMiddleware_ErrorResponse(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	req := httptest.NewRequest(http.MethodGet, "/error", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := CustomLogger()
	wrappedHandler := middleware(handler)

	err := wrappedHandler(c)
	if err == nil {
		t.Error("Expected error, got none")
	}

	// Verificar se é um echo.HTTPError
	if httpErr, ok := err.(*echo.HTTPError); ok {
		if httpErr.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", httpErr.Code)
		}
		if httpErr.Message != "Bad request" {
			t.Errorf("Expected 'Bad request', got '%s'", httpErr.Message)
		}
	} else {
		t.Errorf("Expected echo.HTTPError, got %T", err)
	}
}
