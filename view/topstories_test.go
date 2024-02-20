package view

import (
	"errors"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/mock"
)

func TestTopStoriesView(t *testing.T) {
	t.Run("when the repository returns an error", func(t *testing.T) {
		testerr := errors.New("test")
		repo := mock.TopStoryIdsRepo{Err: testerr}
		view := TopStoriesView{Repo: repo}
		_, err := view.Data(
			mock.Params{},
			mock.Headers{},
			Options{Layout: "part"},
		)
		assert.Equal(t, err, testerr)
	})

	t.Run("when the view renders succesfully", func(t *testing.T) {
		ids := []int{42, 3545, 345, 1}
		repo := mock.TopStoryIdsRepo{Ids: ids}
		view := TopStoriesView{Repo: repo}

		data, err := view.Data(
			mock.Params{},
			mock.Headers{},
			Options{Layout: "part"},
		)

		assert.NoError(t, err)
		assert.Equal(t, data, ViewData[TopStoriesViewData]{
			Template: "topstories.gohtml",
			HashKey:  "",
			Data: TopStoriesViewData{
				Ids: ids,
			},
		})
	})
}
