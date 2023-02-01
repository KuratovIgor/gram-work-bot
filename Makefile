.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t gram-work-bot-image .

start-container:
	docker run --name gram-work-bot -p 80:80 --env-file .env gram-work-bot-image