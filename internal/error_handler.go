package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CustomErrorHandler implementa o tratamento de erros centralizado
func CustomErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Erro interno do servidor"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Message != nil {
			if msg, ok := he.Message.(string); ok {
				message = msg
			}
		}
	}

	// Log do erro
	c.Logger().Error(err)

	// Resposta de erro
	if !c.Response().Committed {
		if c.Request().Header.Get("Content-Type") == "application/xml" {
			if err := c.XML(code, map[string]interface{}{
				"error":   message,
				"code":    code,
				"success": false,
			}); err != nil {
				c.Logger().Error(err)
			}
		} else {
			if err := c.JSON(code, map[string]interface{}{
				"success": false,
				"message": message,
				"error":   "",
			}); err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
