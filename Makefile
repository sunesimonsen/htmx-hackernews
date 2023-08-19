.PHONY: deploy
deploy:
	gcloud app deploy

.PHONY: browse
browse: deploy
	gcloud app browse
