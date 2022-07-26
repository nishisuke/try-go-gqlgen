FROM golang:1.18-bullseye

RUN go install github.com/cosmtrek/air@latest

CMD air
