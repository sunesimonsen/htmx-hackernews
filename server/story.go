package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) Story() Handle {
	type Item struct {
		Id          int    `json:"id"`
		By          string `json:"by"`
		Title       string `json:"title"`
		Url         string `json:"url"`
		Descendants int    `json:"descendants"`
		Type        string `json:"type"`
		Time        int    `json:"time"`
		Score       int    `json:"score"`
		Kids        []int  `json:"kids"`
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

		story := Item{}
		err = json.Unmarshal(storyData, &story)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		return s.templates.ExecuteTemplate(
			w,
			"story.gohtml",
			story,
		)
	}
}
