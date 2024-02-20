package templates

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"text/template"
)

var (
	//go:embed "layouts/*.gohtml" "includes/*.gohtml" "pages/*.gohtml"
	templatesFS embed.FS
)

type Renderer struct {
	templates map[string]*template.Template
}

func NewRenderer() Renderer {
	layoutFiles, err := fs.Glob(templatesFS, "layouts/*.gohtml")
	if err != nil {
		panic(err)
	}

	includeFiles, err := fs.Glob(templatesFS, "includes/*.gohtml")
	if err != nil {
		panic(err)
	}

	pagesFiles, err := fs.Glob(templatesFS, "pages/*.gohtml")
	if err != nil {
		panic(err)
	}

	templates := map[string]*template.Template{}

	for _, file := range pagesFiles {
		fileName := filepath.Base(file)

		files := []string{}
		files = append(files, includeFiles...)
		files = append(files, layoutFiles...)
		files = append(files, file)

		templates[fileName] = template.Must(template.ParseFS(templatesFS, files...))
	}

	return Renderer{templates: templates}
}

func (tr Renderer) Render(writer io.Writer, templateName string, layout string, data any) error {
	template, ok := tr.templates[templateName]

	if !ok {
		return fmt.Errorf("Template doesn't exist: %s", templateName)
	}

	return template.ExecuteTemplate(writer, layout, data)
}
