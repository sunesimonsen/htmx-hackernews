package server

import (
	"testing"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestStory(t *testing.T) {
	story := model.Story{
		By:          "ColinWright",
		Descendants: 112,
		Id:          37173339,
		Kids:        []int{37175324, 37176779, 37175611},
		Score:       191,
		Time:        1692351932,
		Title:       "Short session expiration does not help security",
		Url:         "https://www.sjoerdlangkemper.nl/2023/08/16/session-timeout/",
	}

	t.Run("snapshot GET /story/37173339 with a succesful response", func(t *testing.T) {
		snapshotRequest(t, "/story/37173339", "/v0/item/37173339.json", story)
	})

	t.Run("snapshot GET /parts/story/37173339 with a succesful response", func(t *testing.T) {
		snapshotRequest(t, "/parts/story/37173339", "/v0/item/37173339.json", story)
	})

	t.Run("GET /parts/story/37173339 with a succesful response", func(t *testing.T) {
		response := renderResponse(t, "/parts/story/37173339", "/v0/item/37173339.json", story)
		doc := newDocument(response)
		assertSelectedText(t, doc, ".title", story.Title)
		assertSelectedText(t, doc, ".by-line", "%d points by %s", story.Score, story.By)
	})
}
