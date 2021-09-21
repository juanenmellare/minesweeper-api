FROM golang:1.16-alpine

RUN apk add git gcc g++ curl

ENV GOPATH /go

WORKDIR /go/src/minesweeper-api

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go mod init

ENTRYPOINT /go/bin/air run /main.go

EXPOSE 8080
