ARG GO_VERSION=1.19
#build stage
FROM golang:${GO_VERSION}-alpine AS builder
# gcc needs musl-dev on alpine, alternatively use libc6-compat (not preferred)
RUN apk add --no-cache --update musl-dev
RUN apk add --no-cache --update gcc
# install make and git
RUN apk add --no-cache --update git make
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN make goget
RUN make test
RUN go install -v ./...
CMD make test
