package repo

import (
	"encoding/json"
	"io"
	"net/http"
)

type Host struct {
	url string
}

func NewHost(url string) Host {
	return Host{url: url}
}

func HackerNewsHost() Host {
	return NewHost("https://hacker-news.firebaseio.com")
}

func (host Host) LoadJson(path string, target any) error {
	url := host.url + path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		return HttpError{
			StatusCode: response.StatusCode,
		}
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
