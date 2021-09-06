# tool macros
GO:=go
GOTEST:=$(GO) test
GOFORMAT:=GOFORMAT
GOLINT:=golint

# version macros
DATE := $(shell /bin/date -u +"%Y-%m-%d~%H:%M:%S")
COMMIT := $(shell git describe --always)

# install path macros
INSTALL_PREFIX := /go
DIST_PATH := dist

# src code path macros
SRC_PATH := cmd/titan
SRC_NAME := titan.go
SRC := $(SRC_PATH)/$(SRC_NAME)

# target macros
TARGET_NAME := titan
BIN_PATH := bin
TARGET := $(BIN_PATH)/$(TARGET_NAME)

# clean files list
DIST_CLEAN_LIST := $(DIST_PATH)
CLEAN_LIST := $(TARGET) \
			  $(DIST_CLEAN_LIST)

# non-phony targets
$(TARGET):
	$(GO) build -ldflags "-X main.date=$(DATE) -X main.commit=$(COMMIT)" -o $(TARGET) $(SRC)

# phony rules
.PHONY: all
all: clean lint test makedir $(TARGET)

.PHONY: makedir
makedir:
	@mkdir -p $(BIN_PATH) $(DIST_PATH)

.PHONY: format
format:
	$(GOFORMAT) ./...

.PHONY: lint
lint:
	$(GOLINT) ./...

.PHONY: test
test:
	$(GOTEST) ./...
	$(GOTEST) -bench=. ./...

.PHONY: mod-init
mod-init:
	$(GO) mod init github.com/muzammilar/examples-go/titan-prometheus/titan

.PHONY: mod-deps
mod-deps:
	$(GO) get -d -v ./...

.PHONY: clean
clean:
	@echo CLEAN $(CLEAN_LIST)
	@rm -f $(CLEAN_LIST)

.PHONY:
distclean:
	@echo CLEAN $(DIST_CLEAN_LIST)
	@rm -f $(DIST_CLEAN_LIST)

# usually used for installing packages
.PHONY: install
install: $(TARGET)
	install -v -d $(DESTDIR)$(INSTALL_PREFIX)/bin
	install -v $(TARGET) $(DESTDIR)$(INSTALL_PREFIX)/bin

.PHONY: uninstall
uninstall:
	-rm -rf $(DESTDIR)$(INSTALL_PREFIX)/bin
