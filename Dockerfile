FROM golang:1.11.5-alpine

ENV GOPATH $GOPATH:/go/src
ENV APP_ROOT /go/src/goraphql
WORKDIR $APP_ROOT

COPY . $APP_ROOT

RUN apk update \
  && apk add --no-cache git libc-dev gcc \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure