## Variables
PWD := $(shell pwd)
# number of cpus for benchmark testing (a list can be provided)
BENCH_CPUS := 1,2,4,8,16,32
# number of iterations for benchmark testing - no whitespace after variable
# If comment is added after number below, that would add a space
#BENCH_ITERATIONS :=1000 # extra space added
#BENCH_ITERATIONS :=1000# no extra space added
BENCH_ITERATIONS :=1000000
# Parent Directory
PARENT_PATH:=$(abspath ..)

## Builds
.PHONY: all clean test lint goget

all: clean test

clean:
	echo "Nothing to clean in ${PWD}"

test:
# check for raceconditions
	go test -race ./...
# run the benchmark tests
	go test -cpu ${BENCH_CPUS} -benchmem -benchtime ${BENCH_ITERATIONS}x -bench=. ./...

lint:
	golint ./...

# add module requirements
goget:
	go get -d -v ./...
