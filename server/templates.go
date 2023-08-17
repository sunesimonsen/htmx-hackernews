package server

import (
	"bytes"
	"embed"
	"io"
	"math/rand"
	"text/template"
)

var (
	//go:embed "templates/*"
	templatesFS embed.FS
)

func (s *server) setupTemplates() error {
	templates, err := template.New("templates").Funcs(template.FuncMap{
		"randomInt": func(min int, max int) int {
			return rand.Intn(max-min+1) + min
		},
	}).ParseFS(templatesFS, "templates/*.gohtml")

	s.templates = templates

	return err
}

func (s *server) renderTemplate(w io.Writer, name string, data any) error {
	buf := &bytes.Buffer{}
	buf.Grow(512)

	err := s.templates.ExecuteTemplate(
		buf,
		name,
		data,
	)

	if err != nil {
		return err
	}

	buf.WriteTo(w)
	return nil
}
