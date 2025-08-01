package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func setupTestEcho() *echo.Echo {
	e := echo.New()
	return e
}

func TestHandlers_HomeHandler(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandlers()
	err := h.HomeHandler(c)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["success"] != true {
		t.Error("Expected success to be true")
	}

	if response["message"] != "Bem-vindo ao Echo Playground!" {
		t.Errorf("Unexpected message: %v", response["message"])
	}
}

func TestHandlers_HelloHandler(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/hello/João", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("name")
	c.SetParamValues("João")

	h := NewHandlers()
	err := h.HelloHandler(c)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	expectedResponse := "Olá, João! Bem-vindo ao Echo!"
	if strings.TrimSpace(rec.Body.String()) != expectedResponse {
		t.Errorf("Expected '%s', got '%s'", expectedResponse, strings.TrimSpace(rec.Body.String()))
	}
}

func TestHandlers_XMLHandler(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/xml", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandlers()
	err := h.XMLHandler(c)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/xml") {
		t.Errorf("Expected XML content type, got %s", contentType)
	}
}

func TestHandlers_CreateUserHandler(t *testing.T) {
	e := setupTestEcho()

	userJSON := `{"name":"Test User","email":"test@example.com","age":25}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandlers()
	err := h.CreateUserHandler(c)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["success"] != true {
		t.Error("Expected success to be true")
	}

	if response["message"] != "Usuário criado com sucesso" {
		t.Errorf("Unexpected message: %v", response["message"])
	}
}

func TestHandlers_CreateUserHandler_InvalidJSON(t *testing.T) {
	e := setupTestEcho()

	invalidJSON := `{"name":"Test User","email":}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandlers()
	err := h.CreateUserHandler(c)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rec.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["success"] == true {
		t.Error("Expected success to be false")
	}
}

func TestHandlers_SearchHandler(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/search?q=test&category=framework&limit=5", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandlers()
	err := h.SearchHandler(c)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["success"] != true {
		t.Error("Expected success to be true")
	}

	if response["message"] != "Busca realizada" {
		t.Errorf("Unexpected message: %v", response["message"])
	}
}
