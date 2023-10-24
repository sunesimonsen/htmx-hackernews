deploy:
	gcloud app deploy

browse:
	gcloud app browse

test:
	go test ./...

test-update:
	UPDATE_SNAPS=true go test ./...

cover:
	go test ./... -cover
