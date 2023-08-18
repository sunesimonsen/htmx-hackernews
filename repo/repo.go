package repo

type Host struct {
	url string
}

func HackerNewsHost() Host {
	return Host{url: "https://hacker-news.firebaseio.com"}
}
