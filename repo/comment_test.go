package repo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/sunesimonsen/htmx-hackernews/model"
)

func newTestServer(t *testing.T, path string, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Errorf("Expected to request '%s', got: %s", path, r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(response)
		assert.NoError(t, err)

		w.Write(data)
	}))
}

func TestGetComment(t *testing.T) {
	t.Run("with a succesful response", func(t *testing.T) {
		responseComment := model.Comment{
			Id:   42,
			Text: "<p>Comment body</p>",
		}
		server := newTestServer(t, "/v0/item/42.json", responseComment)
		defer server.Close()

		host := NewHost(server.URL)

		comment, err := host.GetComment("42")
		assert.NoError(t, err)
		assert.Equal(t, comment, responseComment)
	})

	t.Run("with a null response", func(t *testing.T) {
		server := newTestServer(t, "/v0/item/42.json", nil)
		defer server.Close()

		host := NewHost(server.URL)

		_, err := host.GetComment("42")
		assert.EqualError(t, err, "Not found: comment 42")
	})
}
