FROM golang:1.22.2-alpine3.19

WORKDIR /app
# git
RUN apk add --no-cache git

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

RUN apk add --no-cache npm



COPY go.mod go.sum ./
COPY package.json package-lock.json ./

RUN go mod download
RUN npm install

COPY . .
RUN npm run build

EXPOSE 8080
