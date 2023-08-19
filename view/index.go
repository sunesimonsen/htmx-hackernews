package view

type IndexView struct {
	Templates TemplateRenderer
}

func (v IndexView) Render(params Params, headers Headers) ([]byte, error) {
	return v.Templates.Render("index.gohtml", nil)
}
