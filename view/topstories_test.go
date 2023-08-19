package view

import (
	"errors"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

func TestTopStoriesView(t *testing.T) {
	renderer := templates.NewRenderer()

	t.Run("when the repository returns an error", func(t *testing.T) {
		testerr := errors.New("test")
		repo := mock.TopStoryIdsRepo{Err: testerr}
		view := TopStoriesView{Repo: repo, Templates: renderer}
		_, err := view.Render(mock.Params{}, mock.Headers{})
		assert.Equal(t, err, testerr)
	})

	t.Run("when the view renders succesfully", func(t *testing.T) {
		ids := []int{42, 3545, 345, 1}
		repo := mock.TopStoryIdsRepo{Ids: ids}
		view := TopStoriesView{Repo: repo, Templates: renderer}

		data, err := view.Render(mock.Params{}, mock.Headers{})

		assert.NoError(t, err)
		snaps.MatchSnapshot(t, string(data))
	})
}
