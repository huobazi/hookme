
.PHONY: all
all: build release

.PHONY: help
help:
	@echo "Usage:"
	@echo "make - compile the source code"
	@echo "make clean - remove binary file"

.PHONY: build
build: deps
	@bash scripts/build.sh

release: clean deps ## Generate releases for all platforms
	@bash scripts/release.sh


.PHONY: deps
deps: ## Install dependencies using go get
	@go get -d -v -t ./...

.PHONY: clean
clean:
	@bash scripts/clean.sh