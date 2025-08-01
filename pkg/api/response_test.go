package api

import (
	"encoding/json"
	"testing"
)

func TestNewSuccessResponse(t *testing.T) {
	message := "Operação realizada com sucesso"
	data := map[string]interface{}{
		"id":   123,
		"name": "Test",
	}

	response := NewSuccessResponse(message, data)

	if !response.Success {
		t.Error("Expected Success to be true")
	}

	if response.Message != message {
		t.Errorf("Expected message %s, got %s", message, response.Message)
	}

	if response.Data == nil {
		t.Error("Expected data to be set, got nil")
	}
}

func TestNewSuccessMessage(t *testing.T) {
	message := "Operação concluída"

	response := NewSuccessMessage(message)

	if !response.Success {
		t.Error("Expected Success to be true")
	}

	if response.Message != message {
		t.Errorf("Expected message %s, got %s", message, response.Message)
	}

	if response.Data != nil {
		t.Error("Expected Data to be nil")
	}
}

func TestNewErrorResponse(t *testing.T) {
	message := "Erro na operação"
	detail := "Detalhes do erro"

	response := NewErrorResponse(message, detail)

	if response.Success {
		t.Error("Expected Success to be false")
	}

	if response.Message != message {
		t.Errorf("Expected message %s, got %s", message, response.Message)
	}

	if response.Error != detail {
		t.Errorf("Expected error detail %s, got %s", detail, response.Error)
	}
}

func TestResponseJSONSerialization(t *testing.T) {
	tests := []struct {
		name     string
		response *Response
	}{
		{
			"Success response",
			NewSuccessResponse("Success", map[string]string{"key": "value"}),
		},
		{
			"Success message",
			NewSuccessMessage("Simple success"),
		},
		{
			"Error response",
			NewErrorResponse("Error occurred", "Details"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Serializar para JSON
			jsonData, err := json.Marshal(tt.response)
			if err != nil {
				t.Errorf("Failed to marshal response to JSON: %v", err)
			}

			// Deserializar de volta
			var response Response
			err = json.Unmarshal(jsonData, &response)
			if err != nil {
				t.Errorf("Failed to unmarshal JSON to response: %v", err)
			}

			// Verificar se os campos principais foram preservados
			if response.Success != tt.response.Success {
				t.Errorf("Expected Success %t, got %t", tt.response.Success, response.Success)
			}

			if response.Message != tt.response.Message {
				t.Errorf("Expected Message %s, got %s", tt.response.Message, response.Message)
			}
		})
	}
}
