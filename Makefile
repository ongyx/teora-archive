# User-defined variables
DEBUG := 0

MAIN := main/main.go
BUILD := build
BINARY := $(BUILD)/teora
TAGS :=
WFLAGS :=

ifneq ($(DEBUG),0)
	TAGS := debug #ebitendebug
else
	WFLAGS := -ldflags -H=windowsgui
endif

.PHONY: all bootstrap clean

all: native

native:
	go build -o $(BINARY) -tags=$(TAGS) $(MAIN)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows:
	go build -o $(BINARY).exe -tags=$(TAGS) $(WFLAGS) $(MAIN)

bootstrap:
	$(MAKE) -C assets all

clean:
	rm -r $(BUILD)
	$(MAKE) -C assets clean
