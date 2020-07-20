FROM golang:1.14-alpine as builder

ADD . /go/src/pictweet-go
WORKDIR /go/src/pictweet-go

ENV GO111MODULE=on

RUN apk update && \
    apk add git && \
    GOOS=linux GOARCH=amd64 go build cmd/pictweet/main.go

FROM alpine:3.11
WORKDIR /app
COPY --from=0 /go/src/pictweet-go /app

CMD ["./main"]
