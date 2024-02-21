package server

import (
	"testing"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestComment(t *testing.T) {
	comment := model.Comment{
		By:     "CapitalistCartr",
		Id:     37176230,
		Parent: 37175721,
		Text:   "The link seems to be Slashdotted into oblivion.",
		Time:   1692369360,
		Kids:   []int{42, 24},
	}

	t.Run("snapshot GET /comment/37176230 with a succesful response", func(t *testing.T) {
		snapshotRequest(t, "/comment/37176230", "/v0/item/37176230.json", comment)
	})

	t.Run("snapshot GET /parts/comment/37176230 with a succesful response", func(t *testing.T) {
		snapshotRequest(t, "/parts/comment/37176230", "/v0/item/37176230.json", comment)
	})

	t.Run("GET /comment/37176230 with a succesful response", func(t *testing.T) {
		response := renderResponse(t, "/comment/37176230", "/v0/item/37176230.json", comment)
		doc := newDocument(response)
		assertSelectedText(t, doc, ".comment-body", comment.Text)

		assertNotMatching(t, doc, "a[href=\"/comment/37176230\"]")
		assertMatching(t, doc, ".comments")
		assertMatching(t, doc, "article[hx-get=\"/parts/comment/42\"]")
	})

	t.Run("GET /parts/comment/37176230 with a succesful response", func(t *testing.T) {
		response := renderResponse(t, "/parts/comment/37176230", "/v0/item/37176230.json", comment)
		doc := newDocument(response)
		assertSelectedText(t, doc, ".comment-body", comment.Text)

		assertMatching(t, doc, "a[href=\"/comment/37176230\"]")
		assertNotMatching(t, doc, ".comments")
	})
}
