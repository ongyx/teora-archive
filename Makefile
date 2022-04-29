# User-defined variables
DEBUG := 1
PPROF := 0

BUILD := build
BINARY := $(BUILD)/teora
TAGS :=
WFLAGS :=

ifeq ($(DEBUG),1)
	TAGS := debug
else
	WFLAGS := -ldflags -H=windowsgui
endif

ifeq ($(PPROF),1)
	TAGS := $(TAGS),pprof
endif

.PHONY: all bootstrap clean

all: native

native:
	go build -tags=$(TAGS) -o $(BINARY)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows:
	go build -tags=$(TAGS) $(WFLAGS) -o $(BINARY).exe	

bootstrap:
	$(MAKE) -C assets all

clean:
	rm -r $(BUILD)
	$(MAKE) -C assets clean
