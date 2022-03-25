MAIN := main/main.go
BUILD := build
BINARY := $(BUILD)/teora
TAGS :=

.PHONY: all assets clean

all: release

debug: TAGS := debug #ebitendebug
debug: native windows

release: WFLAGS := -ldflags -H=windowsgui
release: native windows

native:
	go build -tags $(TAGS) -o $(BINARY) $(MAIN)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows:
	go build $(WFLAGS) -tags $(TAGS) -o $(BINARY).exe $(MAIN)

assets:
	$(MAKE) -C assets all

clean:
	rm -r $(BUILD)
	$(MAKE) -C assets clean
