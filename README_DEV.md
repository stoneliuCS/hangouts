# Project requirements

## Hangouts API
- [Go](https://go.dev)
  - Hangouts server written in Go.
- [Bun](https://bun.sh)
  - OpenAPI generation in TypeScript through _fluid-oas_
- [TaskFile](https://taskfile.dev)
  - For easy builds.
- [gofmt](https://pkg.go.dev/cmd/gofmt)
  - Go formatter.
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
  - Containerization and integration tests.

## Swift Client
- xcode v15.0+ (Required for IOS development)
- [xcodegen](https://github.com/yonaskolb/XcodeGen)
  - Builds the same xcode project across all repos.
- [xcode-build-server](https://github.com/SolaWing/xcode-build-server)
  - Optional, not needed if developing under xcode.

### On development with Swift
- Sometimes the source kit LSP does not register, it just means you need to buil the project. (Only applies of development outside of xcode.)
