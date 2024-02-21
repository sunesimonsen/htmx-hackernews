package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/htmlformat"
	"github.com/alecthomas/assert/v2"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/sunesimonsen/htmx-hackernews/mock"
	"github.com/sunesimonsen/htmx-hackernews/repo"
)

func renderResponse(t *testing.T, path string, upstreamPath string, responseObject any) *http.Response {
	t.Helper()

	upstream := mock.NewServer(t, upstreamPath, responseObject)

	server, err := NewServer(Config{RepoHost: repo.NewHost(upstream.URL)})
	assert.NoError(t, err)

	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	return response.Result()
}

func snapshotResponse(t *testing.T, response *http.Response) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	err := htmlformat.Fragment(buffer, response.Body)
	assert.NoError(t, err)

	snaps.MatchSnapshot(t,
		response.Header,
		response.StatusCode,
		buffer.String(),
	)
}

func snapshotRequest(t *testing.T, path string, upstreamPath string, responseObject any) {
	t.Helper()

	response := renderResponse(t, path, upstreamPath, responseObject)
	snapshotResponse(t, response)
}

func newDocument(response *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

func assertSelectedText(t *testing.T, doc *goquery.Document, selector string, expected string, args ...any) {
	t.Helper()

	assert.Equal(
		t,
		fmt.Sprintf(expected, args...),
		doc.Find(selector).Text(),
		"selector %s", selector,
	)
}

func assertMatching(t *testing.T, doc *goquery.Document, selector string) {
	t.Helper()

	assert.True(t, doc.Find(selector).Length() > 0, "selector %s not found", selector)
}

func assertNotMatching(t *testing.T, doc *goquery.Document, selector string) {
	t.Helper()

	count := doc.Find(selector).Length()

	assert.False(t, count > 0, "selector %s found %d times", selector, count)
}
