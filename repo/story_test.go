package repo

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func TestGetStory(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		responseStory := model.Story{
			Id:    42,
			Title: "This is the title of the story",
		}
		server := mock.NewServer(t, "/v0/item/42.json", responseStory)
		defer server.Close()

		host := NewHost(server.URL)

		story, err := host.GetStory("42")
		assert.NoError(t, err)
		assert.Equal(t, story, responseStory)
	})

	t.Run("with a null response", func(t *testing.T) {
		server := mock.NewServer(t, "/v0/item/42.json", nil)
		defer server.Close()

		host := NewHost(server.URL)

		_, err := host.GetStory("42")
		assert.EqualError(t, err, "Not found: story 42")
	})

	t.Run("with a HTTP error", func(t *testing.T) {
		server := mock.NewFailingServer(500)
		defer server.Close()

		host := NewHost(server.URL)

		_, err := host.GetStory("42")
		assert.EqualError(t, err, "500 Internal Server Error")
	})
}
