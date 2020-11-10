# CHANGELOG

## v0.2.0 (2020-11-10)

### Added

- feat: add exclude-to, include-to flag

- feat: add log include to revision in git

### Fixed

- fix: empty CHANGELOG title when no new commits

- fix: correct exclude-to, include-to revision when get commits

### Others

- chore: correct exclude-to, include-to flag when log

- docs: add github markdown in comment

- refactor: remove useless return error in git

- docs: add comment for git methods

- refactor: move name, description in cli app to const

- build: update go.mod

- chore: generate CHANGELOG v0.1.0

## v0.1.0 (2020-11-10)

### Added

- feat: only add type of change in CHANGELOG when there is changed

- feat: add --version flag to generate CHANGELOG.md

- feat: write changelog to path

- feat: format commit as markdown item

- feat: remove scope and description in conventional commit

- feat: add markdown generator to generate markdown lines

- feat: add RawHeader in conventional commit

- feat: add conventional commits

- feat: remove author and hash commit

- feat: get commits in path between from and to revision

### Others

- chore: add MIT LICENSE

- refactor: remove body and footers in convention

- refactor: use struct action to split long fn

- chore: add go test, lint in github action

- teat: unit test for new conventional commit

- chore: init go mod with gitignore