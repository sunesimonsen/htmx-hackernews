package templates

import (
	"bytes"
	"embed"
	"text/template"
)

var (
	//go:embed "templates/*"
	templatesFS embed.FS
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer() Renderer {
	templates := template.Must(template.New("templates").ParseFS(templatesFS, "templates/*.gohtml"))

	return Renderer{templates: templates}
}

func (tr Renderer) Render(template string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Grow(512)

	err := tr.templates.ExecuteTemplate(
		buf,
		template,
		data,
	)

	return buf.Bytes(), err
}
