package view

import (
	"bytes"
	"embed"
	"text/template"
)

var (
	//go:embed "templates/*"
	templatesFS embed.FS
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() TemplateRenderer {
	templates, err := template.New("templates").ParseFS(templatesFS, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return TemplateRenderer{templates: templates}
}

func (tr TemplateRenderer) Render(template string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Grow(512)

	err := tr.templates.ExecuteTemplate(
		buf,
		template,
		data,
	)

	return buf.Bytes(), err
}
