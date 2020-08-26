FROM golang:1.14-alpine as builder

ADD . /go/src/pictweet-go
WORKDIR /go/src/pictweet-go

ENV GO111MODULE=on

RUN apk update && \
    apk add git && \
    GOOS=linux GOARCH=amd64 go build cmd/pictweet/main.go

ENV TZ Asia/Tokyo
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/pictweet-go /app

CMD ["./main"]

EXPOSE 8080