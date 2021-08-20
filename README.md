# Changeloguru

[![Go](https://github.com/haunt98/changeloguru/workflows/Go/badge.svg?branch=main)](https://github.com/actions/setup-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/haunt98/changeloguru.svg)](https://pkg.go.dev/github.com/haunt98/changeloguru)
[![codecov](https://codecov.io/gh/haunt98/changeloguru/branch/main/graph/badge.svg?token=ZBG353F0CN)](https://codecov.io/gh/haunt98/changeloguru)

Tool to generate `CHANGELOG.md`, `CHANGELOG.rst` from [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## Install

With Go version `>= 1.16`:

```sh
go install github.com/haunt98/changeloguru/cmd/changeloguru@latest
```

With Go version `< 1.16`:

```sh
GO111module=on go get github.com/haunt98/changeloguru/cmd/changeloguru
```

## Usage

```sh
# Help
changeloguru --help

# Generate changelog v1.0.0
changeloguru generate --version v1.0.0

# Generate changelog v2.0.0 from HEAD to tag v1.0.0
changeloguru generate --to v1.0.0 --version v2.0.0

# Generate changelog in dry run mode (without changing anything)
changeloguru generate --to v1.0.0 --version v2.0.0 --dry-run

# Generate changelog only for scope
changeloguru generate --to v1.0.0 --version v2.0.0 --scope projectA --scope projectB
```

## Thanks

- [Command Line Interface Guidelines](https://clig.dev/)
- [fatih/color](https://github.com/fatih/color)
- [go-git/go-git](https://github.com/go-git/go-git)
- [google/go-cmp](https://github.com/google/go-cmp)
- [sebdah/goldie](https://github.com/sebdah/goldie)
- [stretchr/testify](https://github.com/stretchr/testify)
- [urfave/cli](https://github.com/urfave/cli)
