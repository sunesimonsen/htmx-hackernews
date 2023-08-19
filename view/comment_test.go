package view

import (
	"errors"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/model"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

func TestCommentView(t *testing.T) {
	renderer := templates.NewRenderer()

	t.Run("when the repository returns an error", func(t *testing.T) {
		testerr := errors.New("test")
		repo := mock.CommentRepo{Err: testerr}
		view := CommentView{Repo: repo}
		_, err := view.Render(mock.Params{"id": "42"}, mock.Headers{})
		assert.Equal(t, err, testerr)
	})

	t.Run("when the view renders succesfully", func(t *testing.T) {
		comment := model.Comment{
			By:     "CapitalistCartr",
			Id:     37176230,
			Parent: 37175721,
			Text:   "The link seems to be Slashdotted into oblivion.",
			Time:   1692369360,
		}
		repo := mock.CommentRepo{Comment: comment}
		view := CommentView{Repo: repo, Templates: renderer}

		data, err := view.Render(mock.Params{"id": "37176230"}, mock.Headers{"Hx-Request": "true"})

		assert.NoError(t, err)
		snaps.MatchSnapshot(t, string(data))
	})
}
