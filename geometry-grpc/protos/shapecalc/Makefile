
# For now, we use relative paths
SRC_DIR := .
IMP_DIR := ..
DST_DIR := $(SRC_DIR)

.PHONY: all clean protos

all: clean protos

# Include both current directory and parent directory
# All paths are relative to source
protos:
	protoc --proto_path=$(SRC_DIR) \
			--proto_path=$(IMP_DIR) \
			--go_opt=paths=source_relative \
			--go_out=$(DST_DIR) \
			--go-grpc_opt=paths=source_relative	\
			--go-grpc_out=$(DST_DIR) \
		$(SRC_DIR)/shapecalc.proto

clean:
	-rm -f *.pb.go
