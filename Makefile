.PHONY: help, deps, test, test-coverage, test-coverage-html, test-e2e, lint, fmt, docker, build, full

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null || echo "unknown")

.DEFAULT_GOAL := help
help: ## List targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Get dependencies
	go get -d -v ./...

test: ## Run tests
	go test -short ./...

test-coverage: ## Create a code coverage report
	mkdir -p .cover
	go test -coverprofile .cover/cover.out ./...

test-coverage-html: ## Create a code coverage report in HTML
	mkdir -p .cover
	go test -coverprofile .cover/cover.out ./...
	go tool cover -html .cover/cover.out

test-e2e:
	bats test/e2e.bats

lint: ## Run linters
	golangci-lint run

fmt: ## Fix formatting issues
	gofmt -w .

docker: ## Build docker image
	@echo "Version: $(VERSION)"
	@echo "Build date: $(BUILD_DATE)"
	@echo "Git commit: $(GIT_COMMIT)"
	docker buildx build --pull --progress=plain \
		--build-arg VERSION="$(VERSION)" \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
		-t jdschulze/php-fpm_exporter:$(VERSION) .

build: ## Build binary
	mkdir -p dist
	CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=$(VERSION) -X main.date=$(BUILD_DATE) -X main.commit=$(GIT_COMMIT)" -trimpath -o ./dist/php-fpm_exporter .

full: fmt lint test test-coverage build ## Local build pipeline
