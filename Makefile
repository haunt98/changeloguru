.PHONY: all test test-color coverage coverage-cli coverate-html lint format build clean

all:
	go mod tidy
	$(MAKE) format
	$(MAKE) test-color
	$(MAKE) lint
	$(MAKE) build
	$(MAKE) clean

test:
	go test -race -failfast ./...

test-color:
	go install github.com/haunt98/go-test-color@latest
	go-test-color -race -failfast ./...

coverage:
	go test -coverprofile=coverage.out ./...

coverage-cli:
	$(MAKE) coverage
	go tool cover -func=coverage.out

coverage-html:
	$(MAKE) coverage
	go tool cover -html=coverage.out

lint:
	go fix ./...
	golangci-lint run ./...

format:
	go install github.com/haunt98/gofimports/cmd/gofimports@latest
	go install mvdan.cc/gofumpt@latest
	gofimports -w --company github.com/make-go-great,github.com/haunt98 .
	gofumpt -w -extra .

build:
	$(MAKE) clean
	go build ./cmd/changeloguru

clean:
	rm -rf changeloguru
