#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN apk add make git
RUN make mod-init
RUN make mod-deps
RUN make
RUN make install

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/titan /titanserver
CMD ./titanserver
LABEL Name=titan Version=0.0.1
