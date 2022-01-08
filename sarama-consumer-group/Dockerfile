ARG GO_VERSION=1.17

###############################################
# Builder stage
###############################################

# producer
FROM golang:${GO_VERSION}-alpine AS producerbuilder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY ./producer .
RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...

ARG GO_VERSION

# consumer
FROM golang:${GO_VERSION}-alpine AS consumerbuilder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY ./consumer .
RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...

###############################################
# Runner stage
###############################################

#producer
FROM alpine:latest AS producer
RUN apk --no-cache add ca-certificates
COPY --from=producerbuilder /go/bin/app /saramaproducer
CMD ./saramaproducer

#consumer
FROM alpine:latest AS consumer
RUN apk --no-cache add ca-certificates
COPY --from=consumerbuilder /go/bin/app /saramaconsumer
CMD ./saramaconsumer
