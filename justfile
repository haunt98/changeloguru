all: tidy format test-color lint build clean

tidy:
    go mod tidy

test:
    go test -race -failfast ./...

test-color:
    # go install github.com/haunt98/go-test-color@latest
    go-test-color -race -failfast ./...

test-golden:
    fd "golden" --exec-batch rm
    go test ./internal/changelog -update
    go test ./internal/convention -update

coverage:
    go test -coverprofile=coverage.out ./...

coverage-cli: coverage
    go tool cover -func=coverage.out

coverage-html: coverage
    go tool cover -html=coverage.out

lint:
    golangci-lint run ./...

format:
    # go install github.com/haunt98/gofimports/cmd/gofimports@latest
    # go install mvdan.cc/gofumpt@latest
    gofimports -w --company github.com/make-go-great,github.com/haunt98 .
    gofumpt -w -extra .

build:
    go build ./cmd/changeloguru

clean:
    rm -rf changeloguru
