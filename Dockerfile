# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build main

EXPOSE 8000

#SHELL ["go", "run main"]