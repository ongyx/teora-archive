MAIN := main/main.go
BUILD := build
BINARY := $(BUILD)/teora

.PHONY: all data clean

all: release

debug: native windows

release: WFLAGS := -ldflags -H=windowsgui
release: native windows

native: data
	go build -o $(BINARY) $(MAIN)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows: data
	go build $(WFLAGS) -o $(BINARY).exe $(MAIN)

data:
	$(MAKE) -C data all

clean:
	rm -r $(BUILD)
	$(MAKE) -C data clean
