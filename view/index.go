package view

type IndexView struct{}
type IndexViewData struct{}

func (v IndexView) Data(params Params, headers Headers, opt Options) (ViewData[IndexViewData], error) {
	return ViewData[IndexViewData]{Template: "index.gohtml"}, nil
}
