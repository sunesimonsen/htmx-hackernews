package server

import (
	"testing"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestGetStory(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		snapshotBody(t, "/story/37173339", "/v0/item/37173339.json", model.Story{
			By:          "ColinWright",
			Descendants: 112,
			Id:          37173339,
			Kids:        []int{37175324, 37176779, 37175611},
			Score:       191,
			Time:        1692351932,
			Title:       "Short session expiration does not help security",
			Url:         "https://www.sjoerdlangkemper.nl/2023/08/16/session-timeout/",
		})
	})
}
