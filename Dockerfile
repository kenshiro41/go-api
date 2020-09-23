FROM golang:1.14.4-alpine3.12 AS build

RUN apk update && apk add --no-cache git make musl-dev curl

WORKDIR /go/src/github.com/kenshiro41/go_app
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:3.12 AS prod

WORKDIR /app

RUN apk update --no-cache \
  && apk add --no-cache ca-certificates

RUN update-ca-certificates

COPY --from=build /go/src/github.com/kenshiro41/go_app ./

EXPOSE 7890
CMD ["./app"]