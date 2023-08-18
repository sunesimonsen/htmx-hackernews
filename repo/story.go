package repo

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (host *Host) GetStory(id string) (model.Story, error) {
	story := model.Story{}
	err := host.LoadJson(fmt.Sprintf("/v0/item/%s.json", id), &story)

	// zero is not a valid id
	if story.Id == 0 {
		return story, fmt.Errorf("%w: story %s", NotFoundError, id)
	}

	return story, err
}
