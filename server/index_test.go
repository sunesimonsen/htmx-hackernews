package server

import (
	"testing"
)

func TestIndex(t *testing.T) {
	t.Run("GET / with a succesful response", func(t *testing.T) {
		snapshotRequest(t, "/", "/v0/topstories.json", []int{42, 3545, 345, 1})
	})
}
