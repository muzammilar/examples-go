ARG ALPINE_BUILD_VERSION=latest
ARG GO_VERSION=1.17
# build protobuf
FROM golang:${GO_VERSION}-alpine AS proto_builder
ARG SERVER_VERSION
RUN apk add --no-cache protoc git
