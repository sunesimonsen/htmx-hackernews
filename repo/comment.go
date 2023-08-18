package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (host *Host) GetComment(id string) (model.Comment, error) {
	url := fmt.Sprintf(host.url+"/v0/item/%s.json", id)
	comment := model.Comment{}

	storyResponse, err := http.Get(url)
	if err != nil {
		return comment, err
	}

	storyData, err := io.ReadAll(storyResponse.Body)
	if err != nil && err != io.EOF {
		return comment, err
	}

	err = json.Unmarshal(storyData, &comment)
	if err != nil {
		return comment, err
	}

	comment.Answers = len(comment.Kids)

	return comment, nil
}
