package view

import (
	"errors"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestStoryView(t *testing.T) {
	t.Run("when the repository returns an error", func(t *testing.T) {
		testerr := errors.New("test")
		repo := mock.StoryRepo{Err: testerr}
		view := StoryView{Repo: repo}
		_, err := view.Data(
			mock.Params{"id": "42"},
			mock.Headers{},
			Options{Layout: "part"},
		)
		assert.Equal(t, err, testerr)
	})

	t.Run("when the view renders succesfully", func(t *testing.T) {
		story := model.Story{
			By:          "ColinWright",
			Descendants: 112,
			Id:          37173339,
			Kids:        []int{37175324, 37176779, 37175611},
			Score:       191,
			Time:        1692351932,
			Title:       "Short session expiration does not help security",
			Url:         "https://www.sjoerdlangkemper.nl/2023/08/16/session-timeout/",
		}
		repo := mock.StoryRepo{Story: story}
		view := StoryView{Repo: repo}

		data, err := view.Data(
			mock.Params{"id": "37173339"},
			mock.Headers{"Hx-Request": "true"},
			Options{Layout: "part"},
		)

		assert.NoError(t, err)
		assert.Equal(t, data, ViewData[StoryViewData]{
			Template: "story.gohtml",
			HashKey:  "descendants:112,score:191",
			Data: StoryViewData{
				Story:            story,
				ShowCommentsLink: true,
			},
		})
	})
}
