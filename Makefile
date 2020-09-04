SHELL := /bin/bash
BASEDIR = $(shell pwd)
LDFLAGS = $(shell govvv -flags)

.PHONY: build
build:
	@echo Build hookme with ldflags: $(LDFLAGS)
	@go build  -o dist/hookme -ldflags "$(LDFLAGS)" cmd/hookme/main.go