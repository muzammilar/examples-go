

# Variables
COMMIT := $(shell /usr/bin/git describe --always)
DEFAULT_VERSION := 0.0.1 # the default application version
BENCH_CPUS := 1          # number of cpus for benchmark testing
BENCH_ITERATIONS := 1000 # number of iterations for benchmark testing
GOMODULENAME := "github.com/muzammilar/examples-go/geometry-grpc"

SERVER_VERSION := ${DEFAULT_VERSION}
CLIENT_VERSION := ${DEFAULT_VERSION}

# Applications
GEOMSERVER=geomserver
DATASERVER=dataserver
CLIENT=geomclient

.PHONY: all clean certs gomodule protos certs clean_protos clean_gomodule

all: clean protos certs ${DATASERVER} ${GEOMSERVER} ${CLIENT}

clean:
	-rm -f ${DATASERVER}
	-rm -f ${GEOMETRYSERVER}

clean_protos:
	$(MAKE) clean -C protos

protos: clean_protos
	$(MAKE) $@ -C protos

${DATASERVER}: certs
	go build -ldflags "-X main.version=${SERVER_VERSION} main.commit=${COMMIT}" -o  ./cmd/pprinter/pprinter.go

${GEOMSERVER}: certs
	go build -ldflags "-X main.version=${SERVER_VERSION} main.commit=${COMMIT}" -o  ./cmd/pprinter/pprinter.go

${CLIENT}: certs
	go build -ldflags "-X main.version=${CLIENT_VERSION} main.commit=${COMMIT}" -o  ./cmd/pprinter/pprinter.go

test:
# check for raceconditions
	go test -race ./...
# run the benchmark tests
	go test -cpu ${BENCH_CPUS} -benchmem -benchtime ${BENCH_ITERATIONS}x -bench=. ./...

lint:
	golint ./...

# clean the go module
clean_gomodule:
	-rm -f go.mod go.sum

# initialize module
gomodinit:
	go mod init ${GOMODULENAME}

# add module requirements and sums
gomodtidy:
	go mod tidy

gomodule: clean_gomodule gomodinit gomodtidy

goget:
	go get -d -v ./...

clean_certs:
	-rm -f certs/*

# source: https://github.com/denji/golang-tls
certs: clean_certs
	mkdir -p certs
# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
	openssl ecparam -genkey -name secp384r1 -out certs/server.grpc.key
# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	openssl req -new -x509 -sha256 \
	-key certs/server.grpc.key -out certs/server.grpc.crt -days 3650 \
	-subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=server.grpc"
