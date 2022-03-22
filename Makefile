GO=go
MAIN=main/main.go
BINARY=teora
BINARY_DIR=build

.PHONY: all native windows data

all: native windows

native: data
	$(GO) build -o $(BINARY_DIR)/$(BINARY) $(MAIN)

windows: data
	GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 $(GO) build -ldflags -H=windowsgui -o $(BINARY_DIR)/$(BINARY).exe $(MAIN)

data:
	$(MAKE) -C data all
