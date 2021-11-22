
# For now, we use relative paths
SRC_DIR := .
DST_DIR := $(SRC_DIR)

.PHONY: all clean protos

all: clean protos

protos:
	protoc -I=$(SRC_DIR) --go_out=$(DST_DIR) $(SRC_DIR)/shapecalc.proto

clean:
	-rm -f *.pb.go
