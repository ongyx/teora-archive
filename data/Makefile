.PHONY: all fonts

all: fonts
	go generate

fonts:
	cd fonts && python make.py

clean:
	rm fonts/*.go
	rm -r fonts/svg
