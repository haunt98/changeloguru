# CHANGELOG

## v0.11.0 (2021-1-5)

### Added

- 6bf63f6e31ad03712f078391dd4cb6d23ca448f6 feat(convention): add git commit directly to conventional commit

- fddbbdb6eca3548d8b6ede8420d3244451543a35 feat(git): add hash in Commit

- 5701a6d5c4ed9758cbe058cce261eba539ca16c5 feat(cli): remove default version

- dccea784d0e8e6165d31552f98a70c4c9eac7289 feat(cli): remove unused output filename

- d045e1d06a2d421245f2058f0a1bb7e523c066b4 feat(git): stop with begin, end fn

- efa0a7430529901d8b86784f569733051da36782 feat(git): only use git log, remove logIncludeTo

- 3501cb3846c0f0e8c27905cc80eb59e4763a9d1e feat(cli): args is need, I was wrong

- 7dd3918be260177332fe757598ace765d0f5f289 feat(cli): this program doest not need any args

- 2dd6ffd1eb4f68d242ce5fc1b0087140f3756bcd feat: only use --to to replace --exclude-to, --include-to

- ed8e4704e31012f4e679e68773803aee21880e42 feat(cli): use --debug instead of --verbose

- e921459daf409baecebe5a4395f64e932cd078cd feat(cli): correct description cli

- 84cf09fc24bef1ddf56fcff17a8ce07be6a040bb feat: exit gracefully

### Fixed

- 0ac61252f3558048fbc55246e473d9bc25732c30 fix(cli): correct --debug short alias

### Others

- c68f833998e463de94f91b068e4c82c12a9cb8c1 chore(git): useless type cast

- 6c550e8a3f98637ddd7400018476eb1436341c55 chore(cli): lowercase debug message

- 17d8fb3a3d68fe745968b774de0e19aa30716391 chore(cli): remove default flags

- 8945cabfef98b944f02979015b2cbc77ad542e05 chore(readme): remove --include-to, --exclude-to in guide

- a9dec8f4b9857466e258367787b3111224fc6eb5 chore(readme): add thanks Command Line Interface Guidelines

- 2be6b0c2ccd984ea3057d8e5d0fc7b62da8f8779 chore(cli): FILETYPE is a misc

- d06a27d88dcc3979f0317f0f35e67c28d172d125 refactor(cli): rename output path to real output

- 2bcc16dd5f96a25072c35a6c084ef319838ff262 refactor(cli): rename output-dir to output

- fc7c5943499f706e25247bcc32d854002bd88482 chore(cli): better usage text

- b9a2121f8b42147ef1e6610f62f24c89afed38a3 build: update go.mod

- d4ff9f0c3d93a953745a7cb3e951145d32f6741c chore: bump golangci-lint v1.34 in github action

- 279ddecae0468f2880c7771d65f5a0f7650a115a chore(changelog): generate v0.10.0