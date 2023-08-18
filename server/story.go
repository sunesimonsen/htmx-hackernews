package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (s *server) Story() Handle {
	type Data struct {
		Story            model.Story
		IncludeLayout    bool
		ShowCommentsLink bool
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		story, err := s.repo.GetStory(ps.ByName("id"))

		if err != nil {
			return err
		}

		includeLayout := r.Header["Hx-Request"] == nil

		return s.renderTemplate(w, "story.gohtml", Data{
			Story:            story,
			IncludeLayout:    includeLayout,
			ShowCommentsLink: !includeLayout && len(story.Kids) > 0,
		})
	}
}
