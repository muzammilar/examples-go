ARG GO_VERSION=1.17
# build protobuf
FROM golang:${GO_VERSION}-alpine AS proto_builder
#Versions
ARG GO_VERSION
ARG PROTOC_VERSION=3.18.1
RUN apk add --no-cache git
