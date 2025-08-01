package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware cria um middleware para autenticação JWT
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"message": "Token de autorização não fornecido",
					"error":   "",
				})
			}

			// Aqui você implementaria a validação real do JWT
			// Por simplicidade, apenas verificamos se o token existe
			if token != "Bearer valid-token" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"message": "Token inválido",
					"error":   "",
				})
			}

			return next(c)
		}
	}
}
