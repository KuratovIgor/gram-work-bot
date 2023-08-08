FROM golang:1.19-alpine AS builder

RUN go version

COPY . /github.com/KuratovIgor/gram-work-bot/
WORKDIR /github.com/KuratovIgor/gram-work-bot/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:3.18.3

WORKDIR /root/

COPY --from=0 /github.com/KuratovIgor/gram-work-bot/.bin/bot .
COPY --from=0 /github.com/KuratovIgor/gram-work-bot/configs configs/

EXPOSE 80

CMD ["./bot"]