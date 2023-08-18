package repo

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestGetTopStoryIds(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		responseIds := []int{
			0, 1, 2, 3, 4,
		}
		server := newTestServer(t, "/v0/topstories.json", responseIds)
		defer server.Close()

		host := NewHost(server.URL)

		ids, err := host.GetTopStoryIds()
		assert.NoError(t, err)
		assert.Equal(t, ids, responseIds)
	})
}
