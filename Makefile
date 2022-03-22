MAIN := main/main.go
BUILD := build
BINARY := $(BUILD)/teora

.PHONY: all assets clean

all: release

debug: native windows

release: WFLAGS := -ldflags -H=windowsgui
release: native windows

native: assets
	go build -o $(BINARY) $(MAIN)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows: assets
	go build $(WFLAGS) -o $(BINARY).exe $(MAIN)

assets:
	$(MAKE) -C assets all

clean:
	rm -r $(BUILD)
	$(MAKE) -C assets clean
