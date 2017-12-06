PROJECT=github.com/fulldump/kip
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

.PHONY: test all clean dependencies setup example

all: test

clean:
	rm -fr src/
	rm -fr bin/
	rm -fr pkg/

setup:
	mkdir -p src/$(PROJECT)
	rmdir src/$(PROJECT)
	ln -s ../../.. src/$(PROJECT)
	$(GOCMD) get gopkg.in/mgo.v2
	$(GOCMD) get gopkg.in/check.v1

test:
	$(GOCMD) test $(PROJECT)
