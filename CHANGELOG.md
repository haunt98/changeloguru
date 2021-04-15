# CHANGELOG

## v1.9.0 (2021-4-15)

### Added

- feat: show time of each commit as the default (#12) (2021-4-15)

### Others

- chore(changelog): generate v1.8.0 (2021-4-14)

## v1.8.0 (2021-4-14)

### Added

- feat: add dry run (#10)

- feat: use my own color pkg to wrap fatih/color (#9)

### Others

- chore(reamde): add --dry-run and --scope flag in guide

- refactor: better changelog parser and generator (#11)

- build: update go.mod

- chore(changelog): generate v1.7.0

## v1.7.0 (2021-4-11)

### Others

- chore: only run github action on main branch (#8)

- refactor: use better commands and flags name (#7)

- chore(changelog): generate v1.6.0

## v1.6.0 (2021-3-29)

### Others

- docs: documenting export methods (#6)

- build: bump go-git v5.3.0

- chore: better build with many OS

- chore(readme): remove lgtm alerts badge

- chore: use semver for future 1.16 go version

- chore(changelog): generate v1.5.0

## v1.5.0 (2021-3-22)

### Added

- feat: support filter commit scope (#5)

### Others

- build: update go.mod

- chore: bump go 1.16.2 in github action

- chore: use go 1.16.x in github action

- chore: remove gotip from github action

- chore: remove gotip build as time consuming

- chore: use semver go version in github action

- chore(changelog): generate v1.4.0

## v1.4.0 (2021-3-17)

### Added

- feat: deprecated ioutil

### Others

- chore: bump go 1.16 in Dockerfile

- build: bump go 1.16 in go.mod

- build: update go.mod

- chore(readme): add guide for install with go 1.16

- chore: bump go v1.16, golangci-lint v1.37 in github action

- build: update go.mod

- build: update go.mod

- chore(license): bump 2021

- chore(changelog): generate v1.3.0

## v1.3.0 (2021-1-20)

### Others

- chore(markdown): re-format file

- chore: move fmtErr global var

- chore: typo defaultRepository

- build: update go.mod

- build: update go.mod

- chore(readme): add lgtm badge

- chore(changelog): generate v1.2.0

## v1.2.0 (2021-1-12)

### Fixed

- fix: no newline at end of file

### Others

- test(changelog): unit test for misc type

- test(convention): unit test for misc type

- refactor(convention): replace commit parseHeader with updateType

- chore: update gitignore

- chore: build generally with go tip

- chore(changelog): generate v1.1.0

## v1.1.0 (2021-1-7)

### Added

- feat(cli): remove use of args

### Others

- chore(cli): remove unused log debug

- refactor(cli): change name -> appName

- refactor(cli): replace flags map with directly field

- chore(changelog): generate v1.0.0

## v1.0.0 (2021-1-5)

### Added

- feat(convention): remove directly access conventional commit

- feat(cli): remove default version

- feat(cli): remove unused output filename

- feat(git): stop with begin, end fn

- feat(git): only use git log, remove logIncludeTo

- feat(cli): args is need, I was wrong

- feat(cli): this program doest not need any args

- feat: only use --to to replace --exclude-to, --include-to

- feat(cli): use --debug instead of --verbose

- feat(cli): correct description cli

- feat: exit gracefully

### Fixed

- fix(cli): correct --debug short alias

### Others

- chore(git): useless type cast

- chore(cli): lowercase debug message

- chore(cli): remove default flags

- chore(readme): remove --include-to, --exclude-to in guide

- chore(readme): add thanks Command Line Interface Guidelines

- chore(cli): FILETYPE is a misc

- refactor(cli): rename output path to real output

- refactor(cli): rename output-dir to output

- chore(cli): better usage text

- build: update go.mod

- chore: bump golangci-lint v1.34 in github action

- chore(changelog): generate v0.10.0

## v0.10.0 (2020-12-24)

### Added

- feat: highlight error

### Others

- chore(readme): add fatih/color in thanks

- chore(git): better error return

- build: update go.mod

- chore(changelog): generate v0.9.0

## v0.9.0 (2020-12-18)

### Added

- feat: by default show help if use do nothing

- feat: remove -v as alias for --verbose

- feat: use urfave/cli DefaultText

- feat: add --repository and --output-dir

### Others

- chore: add Thanks in README

- chore: remove markdown ext from LICENSE

- chore: generate CHANGELOG v0.8.0

## v0.8.0 (2020-12-17)

### Added

- feat: make 0.1.0 default version, split getChangelogPath

- feat: add --filename, --filetype flags

### Fixed

- fix(changelog): wrong header for fixed type

### Others

- test(changelog): more test for generate changelog

- test(changelog): re-gen golden data for correct header

- chore: log version verbose

- chore: add placeholder of urfave/cli

- chore: generate CHANGELOG v0.7.0

## v0.7.0 (2020-12-2)

### Added

- feat(convention): support mixed-case for type

### Others

- test(convetion): unit test for mixedcase types

- chore: generate CHANGELOG v0.6.1

## v0.6.1 (2020-12-2)

### Others

- refactor: split get flags, get args from cli

- chore: remove new line between badges

- chore: add badge for pkg go dev in README

- build: update go.mod

- chore: generate CHANGELOG v0.6.0

## v0.6.0 (2020-11-29)

### Added

- feat(changelog): use newly markdown parser and generate

- feat(markdown): double newline when generate

- feat: generate 1 line for markdown

- feat: parse markdown to base syntax guide

- feat: use testify assert

- feat: add markdown parser

### Others

- test(changelog): unit test for changelog with markdown

- refactor(changelog): split get version header

- refactor(changelog): remove magic number

- refactor: rename markdown.Base -> markdown.Node

- refactor(markdown): rename parser -> parse

- build: bump go-cmp v0.5.4

- chore: bump golangci-lint v1.33.0 in github action

- chore: make github action run on pull request

- chore: add build using gotip in github action

- chore: generate CHANGELOG v0.5.0

## v0.5.0 (2020-11-23)

### Others

- build: update go.mod

- build: bump goldie v2.5.3

- build: bump go-cmp v0.5.3

- docs: add usage guide for generate changelog first time

- chore: generate CHANGELOG v0.4.0

## v0.4.0 (2020-11-11)

### Added

- feat(convention): make sure header commit is trimmed space

### Others

- docs: add usage guide in README

- chore: generate CHANGELOG v0.3.0

## v0.3.0 (2020-11-11)

### Others

- build: add Dockerfile

- docs: Referencing the workflow file using the file path does not work if the workflow has a `name`

- docs: add github action badge and install guide in README

- refactor: move main.go to root dir for easy go get, go install

- chore: generate CHANGELOG v0.2.1

## v0.2.1 (2020-11-11)

### Fixed

- fix(changelog): correct get lines and skip generate if no new lines

### Others

- chore: generate CHANGELOG v0.2.0

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
