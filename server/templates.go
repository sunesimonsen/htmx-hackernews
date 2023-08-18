package server

import (
	"bytes"
	"embed"
	"net/http"
	"text/template"

	"github.com/hhsnopek/etag"
)

var (
	//go:embed "templates/*"
	templatesFS embed.FS
)

func (s *server) setupTemplates() error {
	templates, err := template.New("templates").ParseFS(templatesFS, "templates/*.gohtml")

	s.templates = templates

	return err
}

func (s *server) renderTemplate(w http.ResponseWriter, r *http.Request, name string, data any) error {
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

	e := etag.Generate(buf.Bytes(), true)

	ifNoneMatch := r.Header.Get("If-None-Match")
	if ifNoneMatch == e {
		w.WriteHeader(http.StatusNotModified)
		return nil
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=10")
	buf.WriteTo(w)
	return nil
}
