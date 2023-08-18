package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) Comment() Handle {
	type Comment struct {
		Id      int    `json:"id"`
		By      string `json:"by"`
		Parent  int    `json:"parent"`
		Text    string `json:"text"`
		Time    int    `json:"time"`
		Kids    []int  `json:"kids"`
		Answers int
	}

	type Data struct {
		Comment          Comment
		ShowCommentsLink bool
		IncludeLayout    bool
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", ps.ByName("id"))

		commentResponse, err := http.Get(url)
		if err != nil {
			return err
		}

		commentData, err := io.ReadAll(commentResponse.Body)
		if err != nil && err != io.EOF {
			return err
		}

		comment := Comment{}
		err = json.Unmarshal(commentData, &comment)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		includeLayout := r.Header["Hx-Request"] == nil
		comment.Answers = len(comment.Kids)

		return s.renderTemplate(w, "comment.gohtml", Data{
			Comment:          comment,
			IncludeLayout:    includeLayout,
			ShowCommentsLink: !includeLayout && comment.Answers > 0,
		})
	}
}
