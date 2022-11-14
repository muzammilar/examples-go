## Variables
PWD := $(shell pwd)
# number of cpus for benchmark testing
BENCH_CPUS := 1
# number of iterations for benchmark testing - no whitespace after variable
# If comment is added after number below, that would add a space
#BENCH_ITERATIONS :=1000 # extra space added
#BENCH_ITERATIONS :=1000# no extra space added
BENCH_ITERATIONS :=1000
# Parent Directory
PARENT_PATH:=$(abspath ..)

## Builds
IPUPDATER:=ipupdater

.PHONY: all clean test lint gomodule clean_gomodule gomodinit gomodtidy goget

all: clean gomodtidy test ${IPUPDATER}

${IPUPDATER}:
	go build -ldflags "-X main.version=${SERVER_VERSION}" -o ${PARENT_PATH}/${BUILD_DIR}/${FSSERVER} ./cmd/ipupdater.go


clean:
	echo "Nothing to clean in ${PWD}"

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
