package utils

// CustomValidator representa um validador customizado
type CustomValidator struct{}

// Validate implementa a interface de validação
func (cv *CustomValidator) Validate(i interface{}) error {
	// Implementar validação customizada aqui
	// Por simplicidade, retornamos nil
	return nil
}
