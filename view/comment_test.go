package view

import (
	"errors"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestCommentView(t *testing.T) {
	t.Run("when the repository returns an error", func(t *testing.T) {
		testerr := errors.New("test")
		repo := mock.CommentRepo{Err: testerr}
		view := CommentView{Repo: repo}
		_, err := view.Data(
			mock.Params{"id": "42"},
			mock.Headers{},
			Options{Layout: "content"},
		)
		assert.Equal(t, err, testerr)
	})

	t.Run("when the view renders succesfully", func(t *testing.T) {
		comment := model.Comment{
			By:      "CapitalistCartr",
			Id:      37176230,
			Parent:  37175721,
			Text:    "The link seems to be Slashdotted into oblivion.",
			Time:    1692369360,
			Answers: 42,
		}
		repo := mock.CommentRepo{Comment: comment}
		view := CommentView{Repo: repo}

		data, err := view.Data(
			mock.Params{"id": "37176230"},
			mock.Headers{"Hx-Request": "true"},
			Options{Layout: "content"},
		)

		assert.NoError(t, err)
		assert.Equal(t, data, ViewData[CommentViewData]{
			Template: "comment.gohtml",
			HashKey:  "answers:42",
			Data: CommentViewData{
				Comment:          comment,
				ShowCommentsLink: true,
			},
		})
	})
}
