# Repository Guidelines

## Project Structure & Module Organization
The Go entry point lives in `main.go` and wires CLI commands to packages under `cmd/`. Use `cmd/server.go` and `cmd/get.go` as references when adding new commands. Core PHP-FPM scraping logic sits in `phpfpm/`, with `phpfpm_test.go` covering unit behavior. Integration assets and docker-compose recipes are under `test/`, while dashboard assets reside in `grafana/`. Keep new operational scripts alongside peers in those directories to simplify discovery.

## Build, Test, and Development Commands
Run `make deps` once to fetch modules. `make test` executes the short Go suite, and `make test-coverage` (or `make test-coverage-html`) produces coverage reports. `make lint` runs `golangci-lint`, and `make fmt` applies `goimports` formatting. During local development you can target the CLI via `go run ./cmd/server.go --help` or build a Docker image with `docker build -t php-fpm_exporter .`.

## Coding Style & Naming Conventions
Use idiomatic Go style: gofmt/goimports enforce tabs for indentation, grouped imports, and receiver naming. Match exported symbol comments to GoDoc style and keep filenames lowercase with underscores only when necessary (`phpfpm_test.go`). When expanding CLI flags, follow the existing `phpfpm.<feature>` naming scheme and document new environment variables in README tables.

## Testing Guidelines
Unit tests live alongside production code in `_test.go` files; mirror the package structure when adding cases. Prefer table-driven tests for new collectors. End-to-end checks run through `make test-e2e`, invoking the Bats script and Docker Compose stack—ensure Docker is running before invoking. Submit coverage artifacts when contributing larger features and note any gaps in PRs.

## Commit & Pull Request Guidelines
History follows Conventional Commits (`feat`, `fix`, `chore(deps)`, etc.); keep subjects under 72 characters and reference issues in the body. Each PR should describe the change, outline validation commands (e.g., `make test`, `make lint`), and link relevant issues or dashboards. Include screenshots for Grafana updates and note configuration changes impacting Docker or Kubernetes manifests.
