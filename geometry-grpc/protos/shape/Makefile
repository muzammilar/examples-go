.PHONY: all clean protos

# For now, we use relative paths
SRC_DIR := .
DST_DIR := $(SRC_DIR)

all: clean protos

# All paths are relative to source
protos:
	protoc --proto_path=$(SRC_DIR) \
			--go_opt=paths=source_relative \
			--go_out=$(DST_DIR) \
		$(SRC_DIR)/shape.proto

clean:
	-rm -f $(DST_DIR)/*.pb.go
