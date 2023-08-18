package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) TopStories() Handle {
	type Data struct {
		Ids []int
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		ids, err := s.repo.GetTopStoryIds()

		if err != nil {
			return err
		}

		return s.renderTemplate(w, "topstories.gohtml", Data{
			Ids: ids,
		})
	}
}
