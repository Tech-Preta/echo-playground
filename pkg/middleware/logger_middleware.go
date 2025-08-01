package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// CustomLogger cria um middleware para logging customizado
func CustomLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Log antes da requisição
			log.Printf("Iniciando requisição: %s %s", c.Request().Method, c.Request().URL.Path)

			err := next(c)

			// Log após a requisição
			duration := time.Since(start)
			log.Printf("Requisição finalizada: %s %s - Duração: %v", c.Request().Method, c.Request().URL.Path, duration)

			return err
		}
	}
}
