package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) TopStories() Handle {
	type Data struct {
		Ids []int
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		url := "https://hacker-news.firebaseio.com/v0/topstories.json"

		storyResponse, err := http.Get(url)
		if err != nil {
			return err
		}

		topStoriesData, err := io.ReadAll(storyResponse.Body)
		if err != nil && err != io.EOF {
			return err
		}

		ids := []int{}
		err = json.Unmarshal(topStoriesData, &ids)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		return s.renderTemplate(w, "topstories.gohtml", Data{
			Ids: ids,
		})
	}
}
