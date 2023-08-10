PACKAGES = ./...
TARGET := ./target
GOPATH := $(shell go env GOPATH)

.PHONY: tidy
tidy: ## Run tidy files
	@echo "==> Running tidy..."
	go mod tidy
	go fmt $(PACKAGES)

.PHONY: generate
generate: ## Run source code generation
	@echo "==> Generating source files..."
	go generate $(PACKAGES)

.PHONY: init
init: install ## Installing binaries
	@echo "==> Initialising..."
	git submodule update --init --recursive

.PHONY: install
install: ## Install development dependencies
	@echo "==> Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: lint
lint: ## Run linter
	@echo "==> Running linter..."
	$(GOPATH)/bin/golangci-lint run --timeout=5m $(PACKAGES)

.PHONY: check
check: ## Run checks
	@echo "==> Running checks..."
	go mod verify
	go vet -all $(PACKAGES)

.PHONY: build
build: check lint ## Build the binary
	@echo "==> Building..."
	@mkdir -p $(TARGET)
	go test -coverprofile=$(TARGET)/coverage.out $(PACKAGES)
	go tool cover -html=$(TARGET)/coverage.out -o $(TARGET)/coverage.html

.PHONY: run-integration-tests
run-integration-tests: ## Run integration tests
	@echo "==> Running integration tests..."
	./integration-tests/bats/bin/bats integration-tests/placeholder.bats
