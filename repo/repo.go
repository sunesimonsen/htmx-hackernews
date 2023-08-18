package repo

import (
	"encoding/json"
	"io"
	"net/http"
)

type Host struct {
	url string
}

func HackerNewsHost() Host {
	return Host{url: "https://hacker-news.firebaseio.com"}
}

func (host Host) LoadJson(path string, target any) error {
	url := host.url + path

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil && err != io.EOF {
		return err
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}

	return nil
}
