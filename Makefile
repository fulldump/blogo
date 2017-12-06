# Project specific variables
PROJECT = blogo
DESCRIPTION = Blog system in Golang

# --- the rest of the file should not need to be configured ---

# GO env
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

GOBUILD = $(GOCMD) build $(FLAGS)


.PHONY: all
all:	build_one

.PHONY: build_one
build_one:
	$(GOBUILD) -o bin/$(PROJECT) $(PROJECT)

.PHONY: build_all
build_all:
	@# https://golang.org/doc/install/source
	GOARCH=amd64 GOOS=linux   $(GOBUILD) -o bin/$(PROJECT).linux64 $(PROJECT)
	GOARCH=386   GOOS=linux   $(GOBUILD) -o bin/$(PROJECT).linux32 $(PROJECT)
	GOARCH=amd64 GOOS=darwin  $(GOBUILD) -o bin/$(PROJECT).mac64 $(PROJECT)
	GOARCH=386   GOOS=darwin  $(GOBUILD) -o bin/$(PROJECT).mac32 $(PROJECT)
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o bin/$(PROJECT).win64.exe $(PROJECT)
	GOARCH=386   GOOS=windows $(GOBUILD) -o bin/$(PROJECT).win32.exe $(PROJECT)

.PHONY: run
run:
	$(GOCMD) run src/$(PROJECT)/main.go
