FROM golang:1.22.2-alpine3.19

# git
RUN apk add --no-cache git

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080
