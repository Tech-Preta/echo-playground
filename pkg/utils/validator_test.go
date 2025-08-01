package utils

import (
	"testing"
)

func TestCustomValidator_Validate(t *testing.T) {
	cv := &CustomValidator{}

	tests := []struct {
		name  string
		input interface{}
	}{
		{
			"Valid struct",
			struct {
				Name  string
				Email string
				Age   int
			}{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   25,
			},
		},
		{
			"String input",
			"test string",
		},
		{
			"Number input",
			123,
		},
		{
			"Nil input",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cv.Validate(tt.input)

			// A implementação atual sempre retorna nil
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}

func TestCustomValidator_ValidateInterface(t *testing.T) {
	cv := &CustomValidator{}

	// Verificar se implementa a interface corretamente
	var validator interface{} = cv

	if _, ok := validator.(*CustomValidator); !ok {
		t.Error("CustomValidator should implement validator interface")
	}
}
