package server

import (
	"testing"
)

func TestIndex(t *testing.T) {
	t.Run("GET / with a succesful response", func(t *testing.T) {
		snapshotResponse(t, "/", "/v0/noop", nil)
	})
}
