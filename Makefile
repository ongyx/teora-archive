GO=go
MAIN=main/main.go
BINARY=teora
BINARY_DIR=build

.PHONY: build

all: native
native:
	$(GO) build -o $(BINARY_DIR)/$(BINARY) $(MAIN)
windows:
	GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 $(GO) build -o $(BINARY_DIR)/$(BINARY).exe $(MAIN)
