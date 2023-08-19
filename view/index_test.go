package view

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

func TestIndexView(t *testing.T) {
	renderer := templates.NewRenderer()

	t.Run("when the view renders succesfully", func(t *testing.T) {
		view := IndexView{Templates: renderer}

		data, err := view.Render(mock.Params{}, mock.Headers{})

		assert.NoError(t, err)
		snaps.MatchSnapshot(t, string(data))
	})
}
