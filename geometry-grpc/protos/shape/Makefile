.PHONY: all clean protos

# For now, we use relative paths
SRC_DIR := .
DST_DIR := $(SRC_DIR)

all: clean protos

protos:
	protoc -I=$(SRC_DIR) --go_out=$(DST_DIR) $(SRC_DIR)/shape.proto

clean:
	-rm -f $(DST_DIR)/*.pb.go
