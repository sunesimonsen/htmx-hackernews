package server

import "net/http"

func showComments(r *http.Request) bool {
	comments := r.URL.Query().Get("comments")
	if comments == "false" {
		return false
	}
	return true
}
