FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-elk-hiro

WORKDIR /go/src/github.com/moemoe89/go-elk-hiro

COPY . /go/src/github.com/moemoe89/go-elk-hiro

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-elk-hiro
