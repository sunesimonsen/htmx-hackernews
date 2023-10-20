deploy:
	gcloud app deploy

browse:
	gcloud app browse

test:
	go test ./...

cover:
	go test ./... -cover
