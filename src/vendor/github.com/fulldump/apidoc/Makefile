PROJECT=github.com/fulldump/apidoc
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

.PHONY: all
all: test

.PHONY: clean
clean:
	rm -fr src/
	rm -fr bin/
	rm -fr pkg/

.PHONY: setup
setup:
	mkdir -p src/$(PROJECT)
	rm -fr src/$(PROJECT)
	ln -s ../../.. src/$(PROJECT)
	$(GOCMD) get github.com/fulldump/golax
	$(GOCMD) get github.com/fulldump/apitest

.PHONY: test
test:
	$(GOCMD) test $(PROJECT)

.PHONY: coverage
coverage:
	rm -fr coverage
	mkdir -p coverage
	$(GOCMD) list $(PROJECT) > coverage/packages
	@i=a ; \
	while read -r P; do \
		i=a$$i ; \
		$(GOCMD) test ./src/$$P -cover -covermode=count -coverprofile=coverage/$$i.out; \
	done <coverage/packages

	echo "mode: count" > coverage/coverage
	cat coverage/*.out | grep -v "mode: count" >> coverage/coverage
	$(GOCMD) tool cover -html=coverage/coverage
