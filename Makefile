.PHONY: all test coverage-cli coverate-html lint

all: test lint

test:
	go test -race -coverprofile=coverage.out ./...

coverage-cli: test
	go tool cover -func=coverage.out

coverage-html: test
	go tool cover -html=coverage.out

lint:
	golangci-lint run ./...
