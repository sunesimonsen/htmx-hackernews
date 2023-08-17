package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) Story() Handle {
	type Story struct {
		Id          int    `json:"id"`
		By          string `json:"by"`
		Title       string `json:"title"`
		Url         string `json:"url"`
		Descendants int    `json:"descendants"`
		Time        int    `json:"time"`
		Score       int    `json:"score"`
		Kids        []int  `json:"kids"`
	}

	type Data struct {
		Story         Story
		IncludeLayout bool
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", ps.ByName("id"))

		storyResponse, err := http.Get(url)
		if err != nil {
			return err
		}

		storyData, err := io.ReadAll(storyResponse.Body)
		if err != nil && err != io.EOF {
			return err
		}

		story := Story{}
		err = json.Unmarshal(storyData, &story)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		return s.renderTemplate(w, "story.gohtml", Data{
			Story:         story,
			IncludeLayout: r.Header["Hx-Request"] == nil,
		})
	}
}
