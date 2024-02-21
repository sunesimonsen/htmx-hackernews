default: test

templates/*_templ.go: templates/*.templ
	templ generate templates/*.templ

templ: templates/*_templ.go

generate: templ

deploy: generate
	gcloud app deploy

browse:
	gcloud app browse

run: generate
	go run .

test: generate
	go test ./...

test-update: generate
	UPDATE_SNAPS=true go test ./...

cover: generate
	go test ./... -cover

clean:
	rm -rf templates/*_templ.go
