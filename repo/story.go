package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (host *Host) GetStory(id string) (model.Story, error) {
	url := fmt.Sprintf(host.url+"/v0/item/%s.json", id)
	story := model.Story{}

	storyResponse, err := http.Get(url)
	if err != nil {
		return story, err
	}

	storyData, err := io.ReadAll(storyResponse.Body)
	if err != nil && err != io.EOF {
		return story, err
	}

	err = json.Unmarshal(storyData, &story)
	if err != nil {
		return story, err
	}

	return story, nil
}
