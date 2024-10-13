default: test

templates/*_templ.go: templates/*.templ
	templ generate templates/*.templ

templ: templates/*_templ.go

tmp/main: **/*.go templates/*_templ.go
	go build -o tmp/main

build: tmp/main

generate: templ

deploy: generate
	git push dokku main:master

browse:
	open "https://news.sune.one/"

run: generate
	go run .

test: generate
	go test ./...

test-update: generate
	UPDATE_SNAPS=true go test ./...

dev:
	air

cover: generate
	go test ./... -cover

clean:
	rm -rf tmp templates/*_templ.go
