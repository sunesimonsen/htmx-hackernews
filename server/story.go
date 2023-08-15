package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		Score       int    `json:"score"`
		Kids        []int  `json:"kids"`
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		url := "https://hacker-news.firebaseio.com/v0/item/37046770.json"

		storyResponse, err := http.Get(url)
		if err != nil {
			return err
		}

		storyData, err := ioutil.ReadAll(storyResponse.Body)
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
