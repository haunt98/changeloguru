# Changeloguru

[![Go](https://github.com/haunt98/changeloguru/actions/workflows/go.yml/badge.svg)](https://github.com/haunt98/changeloguru/actions/workflows/go.yml)
[![gitleaks](https://github.com/haunt98/changeloguru/actions/workflows/gitleaks.yml/badge.svg)](https://github.com/haunt98/changeloguru/actions/workflows/gitleaks.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/haunt98/changeloguru.svg)](https://pkg.go.dev/github.com/haunt98/changeloguru)
[![Latest Version](https://img.shields.io/github/v/tag/haunt98/changeloguru)](https://github.com/haunt98/changeloguru/tags)

Tool to generate `CHANGELOG.md`, `CHANGELOG.rst` from
[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

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

`--from` is commit in the future aka new.

`--to` is commit in the past aka old.

Go from **future** to the **past**.

```sh
# Help
changeloguru --help

# Generate changelog v1.0.0
changeloguru generate --version v1.0.0

# Generate changelog v2.0.0 from HEAD to tag v1.0.0
changeloguru generate --version v2.0.0 --to-ref v1.0.0

# Generate changelog in dry run mode (without changing anything)
changeloguru generate --version v2.0.0 --to-ref v1.0.0 --dry-run

# Generate changelog in interactive mode (with instruction) and auto push commit, tag
changeloguru generate -i --auto-commit --auto-tag --auto-push
```

## Thanks

- [Command Line Interface Guidelines](https://clig.dev/)
- [fatih/color](https://github.com/fatih/color)
- [go-git/go-git](https://github.com/go-git/go-git)
- [google/go-cmp](https://github.com/google/go-cmp)
- [sebdah/goldie](https://github.com/sebdah/goldie)
- [stretchr/testify](https://github.com/stretchr/testify)
- [urfave/cli](https://github.com/urfave/cli)

Made with [GoLand](https://www.jetbrains.com/go/). Thanks for supporting open
source projects!
