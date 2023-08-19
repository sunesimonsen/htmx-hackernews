package mock

import "github.com/sunesimonsen/htmx-hackernews/model"

type StoryRepo struct {
	Story model.Story
	Err   error
}

func (r StoryRepo) GetStory(id string) (model.Story, error) {
	return r.Story, r.Err
}

type CommentRepo struct {
	Comment model.Comment
	Err     error
}

func (r CommentRepo) GetComment(id string) (model.Comment, error) {
	return r.Comment, r.Err
}

type TopStoryIdsRepo struct {
	Ids []int
	Err error
}

func (r TopStoryIdsRepo) GetTopStoryIds() ([]int, error) {
	return r.Ids, r.Err
}
