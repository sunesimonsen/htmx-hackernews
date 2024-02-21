package server

import (
	"testing"
)

func TestTopStory(t *testing.T) {
	t.Run("GET /parts/topstories with a succesful response", func(t *testing.T) {
		snapshotResponse(t, "/parts/topstories", "/v0/topstories.json", []int{42, 3545, 345, 1})
	})
}
