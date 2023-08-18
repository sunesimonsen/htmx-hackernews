package repo

import (
	"encoding/json"
	"io"
	"net/http"
)

func (host Host) GetTopStoryIds() ([]int, error) {
	url := host.url + "/v0/topstories.json"

	ids := []int{}

	storyResponse, err := http.Get(url)
	if err != nil {
		return ids, err
	}

	topStoriesData, err := io.ReadAll(storyResponse.Body)
	if err != nil && err != io.EOF {
		return ids, err
	}

	err = json.Unmarshal(topStoriesData, &ids)
	if err != nil {
		return ids, err
	}

	return ids, nil
}
