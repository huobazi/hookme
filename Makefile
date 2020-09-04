SHELL := /bin/bash
BASEDIR = $(shell pwd)
LDFLAGS = $(shell govvv -flags)

.PHONY: build
build:
	@go build  -o dist/hookme -ldflags "$(LDFLAGS)" cmd/hookme/main.go