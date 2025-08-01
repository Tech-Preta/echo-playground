package internal

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template representa o renderizador de templates
type Template struct {
	templates *template.Template
}

// NewTemplate cria um novo renderizador de templates
func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
}

// Render implementa a interface de renderização do Echo
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
