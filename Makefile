.PHONY: lint test build

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...

build:
	go build ./...
