# Description: Makefile for Go projects

BIN_NAME := $(shell basename `pwd`)
PACKAGES = ./...
TARGET := ./target
GOPATH := $(shell go env GOPATH)
BIN_VERSION := $(shell echo $$TAG_VERSION)
ifeq ($(strip $(BIN_VERSION)),)
BIN_VERSION := 1.0.0
endif

.PHONY: all
all: clean init build run-integration-tests ## Run all targets

.PHONY: clean
clean: ## Clean the binary
	@echo "==> Cleaning..."
	rm -rf $(TARGET)

.PHONY: init
init: install ## Installing binaries
	@echo "==> Initialising..."
	go version
	git submodule update --init --recursive

.PHONY: build
build: check lint ## Build the binary
	@echo "==> Building..."
	@mkdir -p $(TARGET)
	go test -coverprofile=$(TARGET)/coverage.out $(PACKAGES)
	go tool cover -html=$(TARGET)/coverage.out -o $(TARGET)/coverage.html
	CGO_ENABLED=0 go build -ldflags " \
				-X github.com/martoc/$(BIN_NAME)/cmd.CLIVersion=$(BIN_VERSION) \
				" \
			-o $(TARGET)/$(BIN_NAME) main.go
	chmod 755 $(TARGET)/$(BIN_NAME)

.PHONY: run-integration-tests
run-integration-tests: ## Run integration tests
	@echo "==> Running integration tests..."
	./integration-tests/run.sh

.PHONY: docs
docs: ## Run docs
	@echo "==> Running docs..."
	godoc -http=:6060

.PHONY: tidy
tidy: ## Run tidy files
	@echo "==> Running tidy..."
	go mod tidy
	go fmt $(PACKAGES)
	gofumpt -d .

.PHONY: generate
generate: ## Run source code generation
	@echo "==> Generating source files..."
	go generate $(PACKAGES)

.PHONY: install
install: ## Install development dependencies
	@echo "==> Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0
	go install github.com/golang/mock/mockgen@v1.6.0
	go install mvdan.cc/gofumpt@v0.6.0
	go install golang.org/x/tools/cmd/godoc@v0.12.0

.PHONY: lint
lint: ## Run linter
	@echo "==> Running linter..."
	$(GOPATH)/bin/golangci-lint run --timeout=5m $(PACKAGES)

.PHONY: check
check: ## Run checks
	@echo "==> Running checks..."
	go mod verify
	go vet -all $(PACKAGES)
