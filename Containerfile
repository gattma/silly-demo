FROM golang:1.23-alpine AS build-env
WORKDIR /app
COPY go.mod ./
COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -o /silly-demo github.com/gattma/silly-demo

FROM alpine:3.14

RUN apk update \
    && apk upgrade\
    && apk add --no-cache tzdata curl

WORKDIR /app
COPY --from=build-env /silly-demo .

EXPOSE 8080
CMD [ "./silly-demo" ]