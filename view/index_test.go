package view

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/mock"
)

func TestIndexView(t *testing.T) {
	t.Run("when the view renders succesfully", func(t *testing.T) {
		view := IndexView{}

		data, err := view.Data(
			mock.Params{},
			mock.Headers{},
		)

		assert.NoError(t, err)
		assert.Equal(t, data, ViewData[IndexViewData]{
			Template: "index.gohtml",
			HashKey:  "",
		})
	})
}
