FROM golang:1.13.5-alpine

LABEL maintainer="Kyrielight"

COPY . /tiona/
RUN chmod 700 /tiona/*

ENTRYPOINT ["go", "run", "/tiona/main.go"]

