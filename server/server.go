package server

import (
	"embed"
	"math/rand"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var (
	//go:embed "templates/*"
	templatesFS embed.FS
)

type server struct {
	templates *template.Template
	router    *httprouter.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer() (*server, error) {

	templates, err := template.New("templates").Funcs(template.FuncMap{
		"randomInt": func(min int, max int) int {
			return rand.Intn(max-min+1) + min
		},
	}).ParseFS(templatesFS, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	router := httprouter.New()

	s := &server{templates: templates, router: router}

	s.setupRoutes()

	return s, nil
}
