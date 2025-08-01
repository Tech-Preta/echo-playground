package models

// Product representa um produto no sistema
type Product struct {
	ID          int     `json:"id" xml:"id"`
	Name        string  `json:"name" xml:"name"`
	Price       float64 `json:"price" xml:"price"`
	Description string  `json:"description" xml:"description"`
	Category    string  `json:"category" xml:"category"`
}

// NewProduct cria um novo produto
func NewProduct(name, description, category string, price float64) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		Description: description,
		Category:    category,
	}
}

// SetID define o ID do produto
func (p *Product) SetID(id int) {
	p.ID = id
}
