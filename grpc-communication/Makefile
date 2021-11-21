

# Variables
COMMIT := $(shell /usr/bin/git describe --always)
DEFAULT_VERSION := 0.0.1 # the default application version
BENCH_CPUS := 1          # number of cpus for benchmark testing
BENCH_ITERATIONS := 1000 # number of iterations for benchmark testing
GOMODULENAME := "github.com/muzammilar/examples-go/grpc-communication"

SERVER_VERSION := ${DEFAULT_VERSION}
CLIENT_VERSION := ${DEFAULT_VERSION}

# Applications
SERVER=areaserver
CLIENT=areaclient

.PHONY: all clean certs gomodule protos certs

all: clean gomodule protos certs ${SERVER} ${CLIENT}

clean:
	-rm -f certs/*
	-rm go.mod go.sum
	$(MAKE) $@ -C protos

protos:
	$(MAKE) $@ -C protos

${SERVER}: protos certs
	go build -ldflags "-X main.version=${SERVER_VERSION} main.commit=${COMMIT}" -o  ./cmd/pprinter/pprinter.go

test:
# check for raceconditions
	go test -race ./...
# run the benchmark tests
	go test -cpu ${BENCH_CPUS} -benchmem -benchtime ${BENCH_ITERATIONS}x -bench=. ./...

lint:
	golint ./...

gomodule:
# initialize module
	go mod init ${GOMODULENAME}
# add module requirements and sums
	go mod tidy

goget:
	go get -d -v ./...

# source: https://github.com/denji/golang-tls
certs:
	mkdir -p certs
# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
	openssl ecparam -genkey -name secp384r1 -out certs/server.grpc.key
# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	openssl req -new -x509 -sha256 \
	-key server.grpc.key -out server.grpc.crt -days 3650 \
	-subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=server.grpc"
