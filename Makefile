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
	go build -o $(BINARY) -tags=$(TAGS) $(MAIN)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows:
	go build -o $(BINARY).exe -tags=$(TAGS) $(WFLAGS) $(MAIN)

assets:
	$(MAKE) -C assets all

clean:
	rm -r $(BUILD)
	$(MAKE) -C assets clean
