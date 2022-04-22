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

.PHONY: all assets bento clean

all: native

native:
	go build -o $(BINARY) -tags=$(TAGS) $(MAIN)

windows: export GOOS = windows
windows: export GOARCH = amd64
windows: export CGO_ENABLED = 1
windows: export CC = x86_64-w64-mingw32-gcc
windows:
	go build -o $(BINARY).exe -tags=$(TAGS) $(WFLAGS) $(MAIN)

bootstrap: assets bento

assets:
	$(MAKE) -C assets all

bento:
	cd bento && go generate

clean:
	rm -r $(BUILD)
	$(MAKE) -C assets clean
