package repo

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestGetComment(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		responseComment := model.Comment{
			Id:   42,
			Text: "<p>Comment body</p>",
		}
		server := mock.NewServer(t, "/v0/item/42.json", responseComment)
		defer server.Close()

		host := NewHost(server.URL)

		comment, err := host.GetComment("42")
		assert.NoError(t, err)
		assert.Equal(t, comment, responseComment)
	})

	t.Run("with a null response", func(t *testing.T) {
		server := mock.NewServer(t, "/v0/item/42.json", nil)
		defer server.Close()

		host := NewHost(server.URL)

		_, err := host.GetComment("42")
		assert.EqualError(t, err, "Not found: comment 42")
	})

	t.Run("with a HTTP error", func(t *testing.T) {
		server := mock.NewFailingServer(500)
		defer server.Close()

		host := NewHost(server.URL)

		_, err := host.GetComment("42")
		assert.EqualError(t, err, "500 Internal Server Error")
	})
}
