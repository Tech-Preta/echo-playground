package api

// Response representa uma resposta padrão da API
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewSuccessResponse cria uma resposta de sucesso
func NewSuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse cria uma resposta de erro
func NewErrorResponse(message, error string) *Response {
	return &Response{
		Success: false,
		Message: message,
		Error:   error,
	}
}

// NewSuccessMessage cria uma resposta de sucesso simples
func NewSuccessMessage(message string) *Response {
	return &Response{
		Success: true,
		Message: message,
	}
}
