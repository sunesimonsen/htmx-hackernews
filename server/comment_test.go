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
	}

	t.Run("GET /comment/37176230 with a succesful response", func(t *testing.T) {
		snapshotResponse(t, "/comment/37176230", "/v0/item/37176230.json", comment)
	})

	t.Run("GET /parts/comment/37176230 with a succesful response", func(t *testing.T) {
		snapshotResponse(t, "/parts/comment/37176230", "/v0/item/37176230.json", comment)
	})
}
