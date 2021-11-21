.PHONY: all clean protos

all: clean protos

protos:
	protoc -I=$(SRC_DIR) --go_out=$(DST_DIR) $(SRC_DIR)/addressbook.proto

clean:
	-rm -f *.pb.go
