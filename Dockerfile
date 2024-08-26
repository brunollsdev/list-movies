FROM golang:1.22.6-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN apk update && apk upgrade

RUN apk add build-base

RUN apk add --no-cache sqlite

COPY . .

CMD [ "go", "run", "main.go" ]