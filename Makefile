.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t kuratovia/gram-work-bot-image .

push-image:
	docker push kuratovia/gram-work-bot-image .

start-container:
	docker run --name gram-work-bot -p 80:80 --env-file .env kuratovia/gram-work-bot-image