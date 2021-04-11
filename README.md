# Changeloguru

[![Go](https://github.com/haunt98/changeloguru/workflows/Go/badge.svg?branch=main)](https://github.com/actions/setup-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/haunt98/changeloguru.svg)](https://pkg.go.dev/github.com/haunt98/changeloguru)

Tool to generate `CHANGELOG.md` from [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## Install

With Go version `>= 1.16`:

```sh
go install github.com/haunt98/changeloguru@latest
```

With Go version `< 1.16`:

```sh
GO111module=on go get github.com/haunt98/changeloguru
```

## Usage

```sh
# Help
changeloguru --help

# Generate changelog v1.0.0
changeloguru generate --version v1.0.0

# Generate changelog v2.0.0 from HEAD to tag v1.0.0
changeloguru generate --to v1.0.0 --version v2.0.0
```

## Thanks

- [Command Line Interface Guidelines](https://clig.dev/)
- [fatih/color](https://github.com/fatih/color)
- [go-git/go-git](https://github.com/go-git/go-git)
- [google/go-cmp](https://github.com/google/go-cmp)
- [sebdah/goldie](https://github.com/sebdah/goldie)
- [stretchr/testify](https://github.com/stretchr/testify)
- [urfave/cli](https://github.com/urfave/cli)
