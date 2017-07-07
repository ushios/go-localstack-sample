FROM golang:1.8-alpine
LABEL maintainer "UshioShugo<ushio.s@gmail.com>"

ENV APP_PATH=${GOPATH}/src/github.com/ushios/go-localstack-sample

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

RUN apk update
RUN apk add --no-cache --virtual .goget \
	git && \
	go get -v ./... && \
	apk del .goget

RUN go build ./...
