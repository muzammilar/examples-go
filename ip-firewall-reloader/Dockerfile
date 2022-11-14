#build stage
FROM golang:1-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...
