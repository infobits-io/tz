.PHONY: lint test build coverage bench clean help

## help: Print available targets
help:
	@echo "Available targets:"
	@echo "  lint       Run golangci-lint (includes auto-fix)"
	@echo "  test       Run tests with -v -race"
	@echo "  build      Build the package"
	@echo "  coverage   Run tests with race detector and generate coverage report"
	@echo "  bench      Run benchmarks with memory allocation stats"
	@echo "  clean      Remove generated artifacts"
	@echo "  help       Print this help message"

## lint: Run golangci-lint (includes auto-fix)
lint:
	golangci-lint run ./...

## test: Run tests with -v -race
test:
	go test -v -race ./...

## build: Build the package
build:
	go build ./...

## coverage: Run tests with race detector and generate coverage report
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

## bench: Run benchmarks with memory allocation stats
bench:
	go test -bench=. -benchmem -count=3 ./...

## clean: Remove generated artifacts
clean:
	rm -f coverage.out
