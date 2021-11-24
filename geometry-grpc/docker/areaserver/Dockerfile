ARG ALPINE_BUILD_VERSION=latest
ARG SERVER_VERSION=0.0.1
ARG CLIENT_VERSION=0.0.1
ARG GO_VERSION=1.17

# ARG are reset after every FROM statement, so need redeclaration
ARG GO_VERSION
FROM geometry/protobuilder:latest AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...

#server final stage
FROM alpine:latest
# ARG are reset after every FROM statement, so need redeclaration
ARG SERVER_VERSION
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=grpc-server Version=${SERVER_VERSION}
