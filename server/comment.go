package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (s *server) Comment() Handle {
	type Data struct {
		Comment          model.Comment
		ShowCommentsLink bool
		IncludeLayout    bool
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		comment, err := s.repo.GetComment(ps.ByName("id"))

		if err != nil {
			return err
		}

		includeLayout := r.Header["Hx-Request"] == nil

		return s.renderTemplate(w, "comment.gohtml", Data{
			Comment:          comment,
			IncludeLayout:    includeLayout,
			ShowCommentsLink: !includeLayout && comment.Answers > 0,
		})
	}
}
