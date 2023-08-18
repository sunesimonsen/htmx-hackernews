package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/repo"
	"github.com/tj/assert"
)

func snapshotBody(t *testing.T, path string, upstreamPath string, responseObject any) {
	t.Helper()

	upstream := mock.NewServer(t, upstreamPath, responseObject)

	server, err := NewServer(Config{RepoHost: repo.NewHost(upstream.URL)})
	assert.NoError(t, err)

	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	snaps.MatchSnapshot(t, response.Body.String())
}
