## 
## These variables can be set to customise building teora:
## * BINARY: The path to the built binary.
## * TAGS: The tags to use to build teora:
## 	* debug: Enable debug mode.
## 	* pprof: Start a live pprof server for profiling.
## * FLAGS: Flags to pass to the 'go build' command.
## 

BINARY := build/teora
TAGS :=
FLAGS :=

.PHONY: bootstrap clean

help:       ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

native:     ## Build teora as a native binary.
	go build -tags=$(TAGS) -o $(BINARY)

windows: export GOOS := windows
windows: export GOARCH := amd64
windows: export CGO_ENABLED := 1
windows: export CC := x86_64-w64-mingw32-gcc
windows:    ## Build teora as a Windows console app.
	go build -tags=$(TAGS) $(FLAGS) -o $(BINARY).exe

windowsgui: export FLAGS := $(FLAGS) -ldflags -H=windowsgui
windowsgui: windows
windowsgui: ## Build teora as a Windows GUI app.

bootstrap:  ## Build teora's assets.
	$(MAKE) -C assets all

clean:      ## Cleanup temporary/built files.
	$(MAKE) -C assets clean
