# Project specific variables
PROJECT = blogo
DESCRIPTION = Blog system in Golang

# --- the rest of the file should not need to be configured ---

# GO env
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

COMMIT = $(shell git log -1 --format="%h" 2>/dev/null || echo "0")
VERSION=$(shell git describe --tags --always)
BUILD_DATE = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
COMPILER = $(shell $(GOCMD) version)
FLAGS = -ldflags "\
  -X $(PROJECT)/constants.COMMIT=$(COMMIT) \
  -X $(PROJECT)/constants.VERSION=$(VERSION) \
  -X $(PROJECT)/constants.BUILD_DATE=$(BUILD_DATE) \
  -X '$(PROJECT)/constants.COMPILER=$(COMPILER)' \
  "

GOBUILD = $(GOCMD) build $(FLAGS)

.PHONY: all
all:	build_one

.PHONY: build_one
build_one: statics
	$(GOBUILD) -o bin/$(PROJECT) $(PROJECT)

.PHONY: build_all
build_all: statics
	@# https://golang.org/doc/install/source
	GOARCH=amd64 GOOS=linux   $(GOBUILD) -o bin/$(PROJECT).linux64 $(PROJECT)
	GOARCH=386   GOOS=linux   $(GOBUILD) -o bin/$(PROJECT).linux32 $(PROJECT)
	GOARCH=amd64 GOOS=darwin  $(GOBUILD) -o bin/$(PROJECT).mac64 $(PROJECT)
	GOARCH=386   GOOS=darwin  $(GOBUILD) -o bin/$(PROJECT).mac32 $(PROJECT)
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o bin/$(PROJECT).win64.exe $(PROJECT)
	GOARCH=386   GOOS=windows $(GOBUILD) -o bin/$(PROJECT).win32.exe $(PROJECT)

.PHONY: statics
statics:
	$(GOCMD) run src/genstatic/genstatic.go --dir=src/www/ --package=statics > src/$(PROJECT)/statics/data.go

.PHONY: run
run: statics
	$(GOCMD) run src/$(PROJECT)/main.go
