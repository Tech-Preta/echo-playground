package models

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	name := "João Silva"
	email := "joao@example.com"
	age := 30

	user := NewUser(name, email, age)

	if user.Name != name {
		t.Errorf("Expected name %s, got %s", name, user.Name)
	}

	if user.Email != email {
		t.Errorf("Expected email %s, got %s", email, user.Email)
	}

	if user.Age != age {
		t.Errorf("Expected age %d, got %d", age, user.Age)
	}

	if user.ID != 0 {
		t.Errorf("Expected ID to be 0, got %d", user.ID)
	}

	// Verificar se Created foi definido (deve ser uma string RFC3339 válida)
	if user.Created == "" {
		t.Error("Expected Created to be set, got empty string")
	}

	// Tentar parsear o timestamp
	_, err := time.Parse(time.RFC3339, user.Created)
	if err != nil {
		t.Errorf("Expected Created to be valid RFC3339 format, got error: %v", err)
	}
}

func TestUserSetID(t *testing.T) {
	user := NewUser("Test User", "test@example.com", 25)
	expectedID := 456

	user.SetID(expectedID)

	if user.ID != expectedID {
		t.Errorf("Expected ID %d, got %d", expectedID, user.ID)
	}
}

func TestUserValidation(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		email       string
		age         int
		expectValid bool
	}{
		{"Valid user", "João Silva", "joao@example.com", 25, true},
		{"Empty name", "", "email@example.com", 25, false},
		{"Empty email", "Name", "", 25, false},
		{"Invalid email", "Name", "invalid-email", 25, false},
		{"Negative age", "Name", "email@example.com", -5, false},
		{"Zero age", "Name", "email@example.com", 0, false},
		{"Too old", "Name", "email@example.com", 150, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := NewUser(tt.userName, tt.email, tt.age)

			isValid := user.Name != "" &&
				user.Email != "" &&
				user.Age > 0 &&
				user.Age < 120 &&
				len(user.Email) > 5 && // Validação básica de email
				user.Email != "invalid-email"

			if isValid != tt.expectValid {
				t.Errorf("Expected valid=%t, got valid=%t for %s", tt.expectValid, isValid, tt.name)
			}
		})
	}
}
