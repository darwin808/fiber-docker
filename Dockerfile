FROM golang:1.22.3-alpine

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . . 

RUN go mod tidy
