.PHONY: init
init: lint-deps

.PHONY: lint-deps
lint-deps: ## Install linter dependencies
	@echo "==> Updating linter dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	go install github.com/client9/misspell/cmd/misspell@v0.3.4

.PHONY: lint
lint: ## Run linter
	@echo "==> Running linter..."
	golangci-lint run --timeout=5m

.PHONY: build
build: lint
	go test -v main.go
