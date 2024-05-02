.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=sicon
VERSION=1.0.0
PREFIX=/usr/local

all:
clean:
install:
check:

win64:
	env GOOS='window's GOARCH='amd64' make EXE='.exe' build/sicon.exe


## -- BLOCK:go --
build/sicon$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/sicon
all: build/sicon$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/sicon$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/sicon$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
## -- BLOCK:sh --
install: install-sh
install-sh:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp bin/sicon-example    $(DESTDIR)$(PREFIX)/bin
## -- BLOCK:sh --
