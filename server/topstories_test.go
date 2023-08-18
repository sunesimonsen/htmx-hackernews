package server

import (
	"testing"
)

func TestGetTopStoryIds(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		snapshotBody(t, "/topstories", "/v0/topstories.json", []int{42, 3545, 345, 1})
	})
}
