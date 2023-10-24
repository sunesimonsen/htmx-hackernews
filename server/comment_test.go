package server

import (
	"testing"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestGetComment(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		snapshotResponse(t, "/comment/37176230", "/v0/item/37176230.json", model.Comment{
			By:     "CapitalistCartr",
			Id:     37176230,
			Parent: 37175721,
			Text:   "The link seems to be Slashdotted into oblivion.",
			Time:   1692369360,
		})
	})
}
