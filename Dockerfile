FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-elk-golang

WORKDIR /go/src/github.com/moemoe89/practicing-elk-golang

COPY . /go/src/github.com/moemoe89/practicing-elk-golang

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-elk-golang
