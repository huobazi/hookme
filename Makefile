BINARY = dist/hookme
SHELL := /bin/bash
BASEDIR = $(shell pwd)
LDFLAGS = $(shell govvv -flags)

.PHONY: all
all: build

.PHONY: help
help:
	@echo "Usage:"
	@echo "make - compile the source code"
	@echo "make clean - remove binary file"

.PHONY: build
build:
	@echo "Build hookme with ldflags: $(LDFLAGS)"
	@go build  -o $(BINARY) -ldflags "$(LDFLAGS)" cmd/hookme/main.go


.PHONY: clean
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi