SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

M = $(shell printf "\033[34;1m▶▶▶\033[0m")

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: wire ${BINARY_DIR} ## Compile the code, build Executable File
	$(info $(M) Building App...)
	@$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: test-clean ## Run tests
	$(info $(M) Running Tests..)
	$(GOCMD) test ./... -cover

test-coverage: test-clean ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

test-it: test-clean ## Run integration tests
	$(info $(M) Running integration tests...) @
	$(GOCMD) test -tags integration_test ./...


test-clean: ## Clean the test cache
	$(info $(M) Cleaning test cache...)
	$(GOCMD) clean -testcache

deps: ## Install dependencies
	@echo "Installing dependencies..."
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -v ./...
	$(GOCMD) mod tidy

mod-tidy: ## Tidy Go module
	$(GOCMD) mod tidy

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	$(info $(M) Running Google Wire...)
	@cd internal/infrastructure/server/di && wire

swag: ## Generate swagger docs
	$(info $(M) Running Swag...)
	swag init -g internal/api/handler/user.go -o cmd/api/docs

clean: ## Remove build related files
	$(info $(M) Cleaning...)
	@rm -rf $(BUILD_DIR)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'