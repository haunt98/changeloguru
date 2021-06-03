=========
CHANGELOG
=========

v0.1.0 (2021-06-03)
===================

Added
-----

- feat: add rst changelog parse (2021-06-03)

- feat: add rst changelog generate (2021-06-03)

- feat: split double newline for markdown (2021-05-31)

- feat: remove markdown.NewLine (2021-05-31)

- feat: return error when commit is empty (2021-05-11)

- Revert "feat: skip empty commit" (2021-05-11)

- feat: skip empty commit (2021-05-11)

- feat: use markdown-go (#19) (2021-05-05)

- feat: add color to dry run (#14) (2021-04-23)

- feat: show time of each commit as the default (#12) (2021-04-15)

- feat: add dry run (#10) (2021-04-14)

- feat: use my own color pkg to wrap fatih/color (#9) (2021-04-14)

- feat: support filter commit scope (#5) (2021-03-22)

- feat: deprecated ioutil (2021-03-17)

- feat(cli): remove use of args (2021-01-07)

- feat(convention): remove directly access conventional commit (2021-01-05)

- feat(cli): remove default version (2021-01-05)

- feat(cli): remove unused output filename (2021-01-05)

- feat(git): stop with begin, end fn (2021-01-05)

- feat(git): only use git log, remove logIncludeTo (2021-01-05)

- feat(cli): args is need, I was wrong (2021-01-05)

- feat(cli): this program doest not need any args (2021-01-05)

- feat: only use --to to replace --exclude-to, --include-to (2021-01-05)

- feat(cli): use --debug instead of --verbose (2021-01-05)

- feat(cli): correct description cli (2021-01-05)

- feat: exit gracefully (2020-12-29)

- feat: highlight error (2020-12-24)

- feat: by default show help if use do nothing (2020-12-18)

- feat: remove -v as alias for --verbose (2020-12-18)

- feat: use urfave/cli DefaultText (2020-12-18)

- feat: add --repository and --output-dir (2020-12-18)

- feat: make 0.1.0 default version, split getChangelogPath (2020-12-17)

- feat: add --filename, --filetype flags (2020-12-17)

- feat(convention): support mixed-case for type (2020-12-02)

- feat(changelog): use newly markdown parser and generate (2020-11-28)

- feat(markdown): double newline when generate (2020-11-27)

- feat: generate 1 line for markdown (2020-11-27)

- feat: parse markdown to base syntax guide (2020-11-27)

- feat: use testify assert (2020-11-27)

- feat: add markdown parser (2020-11-27)

- feat(convention): make sure header commit is trimmed space (2020-11-11)

- feat: add exclude-to, include-to flag (2020-11-10)

- feat: add log include to revision in git (2020-11-10)

- feat: only add type of change in CHANGELOG when there is changed (2020-11-10)

- feat: add --version flag to generate CHANGELOG.md (2020-11-10)

- feat: write changelog to path (2020-11-10)

- feat: format commit as markdown item (2020-11-10)

- feat: remove scope and description in conventional commit (2020-11-10)

- feat: add markdown generator to generate markdown lines (2020-11-10)

- feat: add RawHeader in conventional commit (2020-11-10)

- feat: add conventional commits (2020-11-09)

- feat: remove author and hash commit (2020-11-09)

- feat: get commits in path between from and to revision (2020-11-07)

Fixed
-----

- fix: correct time layout (#13) (2021-04-15)

- fix: no newline at end of file (2021-01-12)

- fix(cli): correct --debug short alias (2021-01-05)

- fix(changelog): wrong header for fixed type (2020-12-17)

- fix(changelog): correct get lines and skip generate if no new lines (2020-11-11)

- fix: empty CHANGELOG title when no new commits (2020-11-10)

- fix: correct exclude-to, include-to revision when get commits (2020-11-10)

Others
------

- build: update go.mod (2021-06-01)

- build: update go.mod (2021-05-30)

- build: update go.mod (2021-05-24)

- chore(readme): correct link how to install (2021-05-11)

- chore(changelog): generate v1.12.0 (2021-05-11)

- refactor: move left, rightScope char to const (2021-05-11)

- docs: explain why skip error commit (2021-05-11)

- chore: no need to check flag when use a.log() (2021-05-11)

- build: update go.mod (2021-05-10)

- chore: remove Dockerfile (2021-05-10)

- build: update go.mod (2021-05-07)

- build: update go.mod (2021-05-05)

- chore(changelog): generate v1.11.0 (2021-05-05)

- refactor: move main to cmd (#18) (2021-05-05)

- refactor: use internal instead pkg (#17) (2021-05-05)

- chore: improve wording, typo (#16) (2021-04-26)

- chore(changelog): generate v1.10.0 (2021-04-23)

- refactor(cli): move all cli related to cli pkg (#15) (2021-04-23)

- chore(changelog): generate v1.9.1 (2021-04-15)

- chore(changelog): generate v1.9.0 (2021-04-15)

- chore(changelog): generate v1.8.0 (2021-04-14)

- chore(reamde): add --dry-run and --scope flag in guide (2021-04-14)

- refactor: better changelog parser and generator (#11) (2021-04-14)

- build: update go.mod (2021-04-14)

- chore(changelog): generate v1.7.0 (2021-04-11)

- chore: only run github action on main branch (#8) (2021-04-11)

- refactor: use better commands and flags name (#7) (2021-04-11)

- chore(changelog): generate v1.6.0 (2021-03-29)

- docs: documenting export methods (#6) (2021-03-29)

- build: bump go-git v5.3.0 (2021-03-29)

- chore: better build with many OS (2021-03-29)

- chore(readme): remove lgtm alerts badge (2021-03-29)

- chore: use semver for future 1.16 go version (2021-03-29)

- chore(changelog): generate v1.5.0 (2021-03-22)

- build: update go.mod (2021-03-18)

- chore: bump go 1.16.2 in github action (2021-03-18)

- chore: use go 1.16.x in github action (2021-03-18)

- chore: remove gotip from github action (2021-03-18)

- chore: remove gotip build as time consuming (2021-03-18)

- chore: use semver go version in github action (2021-03-18)

- chore(changelog): generate v1.4.0 (2021-03-17)

- chore: bump go 1.16 in Dockerfile (2021-03-17)

- build: bump go 1.16 in go.mod (2021-03-17)

- build: update go.mod (2021-03-15)

- chore(readme): add guide for install with go 1.16 (2021-03-15)

- chore: bump go v1.16, golangci-lint v1.37 in github action (2021-03-04)

- build: update go.mod (2021-03-04)

- build: update go.mod (2021-02-19)

- chore(license): bump 2021 (2021-01-21)

- chore(changelog): generate v1.3.0 (2021-01-20)

- chore(markdown): re-format file (2021-01-20)

- chore: move fmtErr global var (2021-01-20)

- chore: typo defaultRepository (2021-01-20)

- build: update go.mod (2021-01-20)

- build: update go.mod (2021-01-18)

- chore(readme): add lgtm badge (2021-01-12)

- chore(changelog): generate v1.2.0 (2021-01-12)

- test(changelog): unit test for misc type (2021-01-12)

- test(convention): unit test for misc type (2021-01-12)

- refactor(convention): replace commit parseHeader with updateType (2021-01-12)

- chore: update gitignore (2021-01-12)

- chore: build generally with go tip (2021-01-07)

- chore(changelog): generate v1.1.0 (2021-01-07)

- chore(cli): remove unused log debug (2021-01-07)

- refactor(cli): change name -> appName (2021-01-07)

- refactor(cli): replace flags map with directly field (2021-01-07)

- chore(changelog): generate v1.0.0 (2021-01-05)

- chore(git): useless type cast (2021-01-05)

- chore(cli): lowercase debug message (2021-01-05)

- chore(cli): remove default flags (2021-01-05)

- chore(readme): remove --include-to, --exclude-to in guide (2021-01-05)

- chore(readme): add thanks Command Line Interface Guidelines (2021-01-05)

- chore(cli): FILETYPE is a misc (2021-01-05)

- refactor(cli): rename output path to real output (2021-01-05)

- refactor(cli): rename output-dir to output (2021-01-05)

- chore(cli): better usage text (2021-01-05)

- build: update go.mod (2021-01-04)

- chore: bump golangci-lint v1.34 in github action (2020-12-31)

- chore(changelog): generate v0.10.0 (2020-12-24)

- chore(readme): add fatih/color in thanks (2020-12-24)

- chore(git): better error return (2020-12-23)

- build: update go.mod (2020-12-18)

- chore(changelog): generate v0.9.0 (2020-12-18)

- chore: add Thanks in README (2020-12-18)

- chore: remove markdown ext from LICENSE (2020-12-18)

- chore: generate CHANGELOG v0.8.0 (2020-12-17)

- test(changelog): more test for generate changelog (2020-12-17)

- test(changelog): re-gen golden data for correct header (2020-12-17)

- chore: log version verbose (2020-12-17)

- chore: add placeholder of urfave/cli (2020-12-17)

- chore: generate CHANGELOG v0.7.0 (2020-12-02)

- test(convetion): unit test for mixedcase types (2020-12-02)

- chore: generate CHANGELOG v0.6.1 (2020-12-02)

- refactor: split get flags, get args from cli (2020-11-30)

- chore: remove new line between badges (2020-11-29)

- chore: add badge for pkg go dev in README (2020-11-29)

- build: update go.mod (2020-11-29)

- chore: generate CHANGELOG v0.6.0 (2020-11-29)

- test(changelog): unit test for changelog with markdown (2020-11-29)

- refactor(changelog): split get version header (2020-11-29)

- refactor(changelog): remove magic number (2020-11-28)

- refactor: rename markdown.Base -> markdown.Node (2020-11-28)

- refactor(markdown): rename parser -> parse (2020-11-28)

- build: bump go-cmp v0.5.4 (2020-11-25)

- chore: bump golangci-lint v1.33.0 in github action (2020-11-23)

- chore: make github action run on pull request (2020-11-23)

- chore: add build using gotip in github action (2020-11-23)

- chore: generate CHANGELOG v0.5.0 (2020-11-23)

- build: update go.mod (2020-11-23)

- build: bump goldie v2.5.3 (2020-11-23)

- build: bump go-cmp v0.5.3 (2020-11-23)

- docs: add usage guide for generate changelog first time (2020-11-11)

- chore: generate CHANGELOG v0.4.0 (2020-11-11)

- docs: add usage guide in README (2020-11-11)

- chore: generate CHANGELOG v0.3.0 (2020-11-11)

- build: add Dockerfile (2020-11-11)

- docs: Referencing the workflow file using the file path does not work if the workflow has a `name` (2020-11-11)

- docs: add github action badge and install guide in README (2020-11-11)

- refactor: move main.go to root dir for easy go get, go install (2020-11-11)

- chore: generate CHANGELOG v0.2.1 (2020-11-11)

- chore: generate CHANGELOG v0.2.0 (2020-11-10)

- chore: correct exclude-to, include-to flag when log (2020-11-10)

- docs: add github markdown in comment (2020-11-10)

- refactor: remove useless return error in git (2020-11-10)

- docs: add comment for git methods (2020-11-10)

- refactor: move name, description in cli app to const (2020-11-10)

- build: update go.mod (2020-11-10)

- chore: generate CHANGELOG v0.1.0 (2020-11-10)

- chore: add MIT LICENSE (2020-11-10)

- refactor: remove body and footers in convention (2020-11-10)

- refactor: use struct action to split long fn (2020-11-10)

- chore: add go test, lint in github action (2020-11-09)

- teat: unit test for new conventional commit (2020-11-09)

- chore: init go mod with gitignore (2020-10-06)
