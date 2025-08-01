package models

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	name := "Laptop Test"
	description := "Laptop para testes"
	category := "Eletrônicos"
	price := 2999.99

	product := NewProduct(name, description, category, price)

	if product.Name != name {
		t.Errorf("Expected name %s, got %s", name, product.Name)
	}

	if product.Description != description {
		t.Errorf("Expected description %s, got %s", description, product.Description)
	}

	if product.Category != category {
		t.Errorf("Expected category %s, got %s", category, product.Category)
	}

	if product.Price != price {
		t.Errorf("Expected price %.2f, got %.2f", price, product.Price)
	}

	if product.ID != 0 {
		t.Errorf("Expected ID to be 0, got %d", product.ID)
	}
}

func TestProductSetID(t *testing.T) {
	product := NewProduct("Test Product", "Description", "Category", 100.0)
	expectedID := 123

	product.SetID(expectedID)

	if product.ID != expectedID {
		t.Errorf("Expected ID %d, got %d", expectedID, product.ID)
	}
}

func TestProductValidation(t *testing.T) {
	tests := []struct {
		name        string
		productName string
		description string
		category    string
		price       float64
		expectValid bool
	}{
		{"Valid product", "Laptop", "Gaming laptop", "Electronics", 1999.99, true},
		{"Empty name", "", "Description", "Category", 100.0, false},
		{"Empty description", "Product", "", "Category", 100.0, false},
		{"Empty category", "Product", "Description", "", 100.0, false},
		{"Negative price", "Product", "Description", "Category", -10.0, false},
		{"Zero price", "Product", "Description", "Category", 0.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := NewProduct(tt.productName, tt.description, tt.category, tt.price)

			isValid := product.Name != "" &&
				product.Description != "" &&
				product.Category != "" &&
				product.Price > 0

			if isValid != tt.expectValid {
				t.Errorf("Expected valid=%t, got valid=%t for %s", tt.expectValid, isValid, tt.name)
			}
		})
	}
}
