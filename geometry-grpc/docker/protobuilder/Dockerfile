ARG GO_VERSION=1.17
# build protobuf
FROM golang:${GO_VERSION}-alpine AS proto_builder
#Versions
ARG GO_VERSION
ARG PROTOC_GEN_GO_VERSION=1.26
ARG PROTOC_GEN_GO_GRPC_VERSION=1.1
# Install modules
RUN apk add --no-cache git protoc make
# Install gRPC go packages
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}}
